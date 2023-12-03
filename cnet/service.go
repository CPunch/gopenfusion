package cnet

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"sync"

	"github.com/CPunch/gopenfusion/cnet/protocol"
	"github.com/CPunch/gopenfusion/internal/config"
)

type PacketHandler func(peer *Peer, pkt protocol.Packet) error

func StubbedPacket(_ *Peer, _ protocol.Packet) error {
	return nil
}

type Service struct {
	listener       net.Listener
	port           int
	Name           string
	ctx            context.Context
	started        chan struct{}
	stopped        chan struct{}
	packetHandlers map[uint32]PacketHandler
	peers          map[chan *PacketEvent]*Peer
	stateLock      sync.Mutex

	// OnDisconnect is called when a peer disconnects from the service.
	// uData is the stored value of the key/value pair in the peer map.
	// It may not be set while the service is running. (eg. srvc.Start() has been called)
	OnDisconnect func(peer *Peer)

	// OnConnect is called when a peer connects to the service.
	// return value is used as the value in the peer map.
	// It may not be set while the service is running. (eg. srvc.Start() has been called)
	OnConnect func(peer *Peer)
}

func RandomPort() (int, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer l.Close()

	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(port)
}

func NewService(ctx context.Context, name string, port int) *Service {
	srvc := &Service{
		port: port,
		Name: name,
	}

	srvc.Reset(ctx)
	return srvc
}

func (srvc *Service) Reset(ctx context.Context) {
	srvc.ctx = ctx
	srvc.packetHandlers = make(map[uint32]PacketHandler)
	srvc.peers = make(map[chan *PacketEvent]*Peer)
	srvc.started = make(chan struct{})
	srvc.stopped = make(chan struct{})
}

// may not be called while the service is running (eg. srvc.Start() has been called)
func (srvc *Service) AddPacketHandler(pktID uint32, handler PacketHandler) {
	srvc.packetHandlers[pktID] = handler
}

type newPeerConnection struct {
	peer    *Peer
	channel chan *PacketEvent
}

func (srvc *Service) Start() error {
	peerConnections := make(chan newPeerConnection)
	defer close(peerConnections)
	go srvc.handleEvents(peerConnections)

	// open listener socket
	var err error
	srvc.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", srvc.port))
	if err != nil {
		return err
	}
	defer srvc.listener.Close()

	log.Printf("%s service hosted on %s:%d\n", srvc.Name, config.GetAnnounceIP(), srvc.port)

	close(srvc.started) // signal that the service has started
	for {
		conn, err := srvc.listener.Accept()
		if err != nil {
			fmt.Println(err)
			// we expect this to happen when the service is stopped
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}

		// create a new peer and pass it to the event loop
		peer := NewPeer(srvc.ctx, conn)
		eRecv := make(chan *PacketEvent)
		peerConnections <- newPeerConnection{channel: eRecv, peer: peer}
		go peer.Handler(eRecv)
	}
}

func (srvc *Service) getPeer(channel chan *PacketEvent) *Peer {
	return srvc.peers[channel]
}

func (srvc *Service) setPeer(channel chan *PacketEvent, peer *Peer) {
	srvc.peers[channel] = peer
}

func (srvc *Service) removePeer(channel chan *PacketEvent) {
	delete(srvc.peers, channel)
}

// returns a channel that is closed when the service has started.
// this is useful if you need to wait until after the service has started.
func (srvc *Service) Started() <-chan struct{} {
	return srvc.started
}

// returns a channel that is closed when the service has stopped.
// this is useful if you need wait until after the service has stopped.
func (srvc *Service) Stopped() <-chan struct{} {
	return srvc.stopped
}

// calls f for each peer in the service passing the peer and the stored uData.
// if f returns false, the iteration is stopped.
// NOTE: the peer map is not locked while iterating, if you're calling this
// outside of the service's event loop, you'll need to lock the peer map yourself.
func (srvc *Service) RangePeers(f func(peer *Peer) bool) {
	for _, peer := range srvc.peers {
		if !f(peer) {
			break
		}
	}
}

// locks the peer map.
func (srvc *Service) Lock() {
	srvc.stateLock.Lock()
}

// unlocks the peer map.
func (srvc *Service) Unlock() {
	srvc.stateLock.Unlock()
}

func (srvc *Service) stop() {
	// OnDisconnect handler might need to do something important
	srvc.RangePeers(func(peer *Peer) bool {
		peer.Kill()
		if srvc.OnDisconnect != nil {
			srvc.OnDisconnect(peer)
		}
		return true
	})

	log.Printf("%s service stopped\n", srvc.Name)
	close(srvc.stopped)
}

// handleEvents is the main event loop for the service.
// it handles all events from the peers and calls the appropriate handlers.
func (srvc *Service) handleEvents(peerPipe <-chan newPeerConnection) {
	defer srvc.stop()

	poll := make([]reflect.SelectCase, 0, 4)

	// add the stop channel and the peer connection channel to our poll queue
	poll = append(poll, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(srvc.ctx.Done()),
	})
	poll = append(poll, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(peerPipe),
	})

	addPoll := func(channel chan *PacketEvent) {
		poll = append(poll, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(channel),
		})
	}

	removePoll := func(index int) {
		poll = append(poll[:index], poll[index+1:]...)
	}

	for {
		chosen, value, recvOK := reflect.Select(poll)
		switch chosen {
		case 0: // cancel signal received, stop the service
			return
		case 1: // new peer, add it to our poll queue
			if !recvOK {
				return
			}

			evnt := value.Interface().(newPeerConnection)
			addPoll(evnt.channel)
			srvc.connect(evnt.channel, evnt.peer)
		default: // peer event
			channel := poll[chosen].Chan.Interface().(chan *PacketEvent)
			peer := srvc.getPeer(channel)
			if peer == nil {
				log.Printf("Unknown peer event: %v", value)
				removePoll(chosen)
				continue
			}

			evnt, ok := value.Interface().(*PacketEvent)
			if !recvOK || !ok || evnt == nil {
				// peer disconnected, remove it from our poll queue
				removePoll(chosen)
				srvc.disconnect(channel, peer)
				continue
			}

			srvc.Lock()
			if err := srvc.handlePacket(peer, evnt.PktID, protocol.NewPacket(evnt.Pkt)); err != nil {
				log.Printf("Error handling packet: %v", err)
				peer.Kill()
			}
			srvc.Unlock()

			// the packet buffer is given to us by the event, so we'll need to make sure to return it to the pool
			protocol.PutBuffer(evnt.Pkt)
		}
	}
}

func (srvc *Service) handlePacket(peer *Peer, typeID uint32, pkt protocol.Packet) error {
	if hndlr, ok := srvc.packetHandlers[typeID]; ok {
		// fmt.Printf("Handling packet %x\n", typeID)
		if err := hndlr(peer, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] unknown packet ID: %x\n", typeID)
	}

	return nil
}

func (srvc *Service) disconnect(channel chan *PacketEvent, peer *Peer) {
	log.Printf("Peer %p disconnected from %s\n", peer, srvc.Name)
	if srvc.OnDisconnect != nil {
		srvc.OnDisconnect(peer)
	}

	srvc.removePeer(channel)
}

func (srvc *Service) connect(channel chan *PacketEvent, peer *Peer) {
	log.Printf("New peer %p connected to %s\n", peer, srvc.Name)
	if srvc.OnConnect != nil {
		srvc.OnConnect(peer)
	}

	srvc.setPeer(channel, peer)
}
