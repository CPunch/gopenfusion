package service

import (
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"sync"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/internal/protocol"
)

type PacketHandler func(peer *protocol.CNPeer, uData interface{}, pkt protocol.Packet) error

func StubbedPacket(_ *protocol.CNPeer, _ interface{}, _ protocol.Packet) error {
	return nil
}

type Service struct {
	listener       net.Listener
	port           int
	Name           string
	stop           chan struct{} // tell active handleEvents() to stop
	stopped        chan struct{}
	started        chan struct{}
	packetHandlers map[uint32]PacketHandler
	peers          map[*protocol.CNPeer]interface{}
	stateLock      sync.Mutex

	// OnDisconnect is called when a peer disconnects from the service.
	// uData is the stored value of the key/value pair in the peer map.
	// It may not be set while the service is running. (eg. srvc.Start() has been called)
	OnDisconnect func(peer *protocol.CNPeer, uData interface{})

	// OnConnect is called when a peer connects to the service.
	// return value is used as the value in the peer map.
	// It may not be set while the service is running. (eg. srvc.Start() has been called)
	OnConnect func(peer *protocol.CNPeer) (uData interface{})
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

func NewService(name string, port int) *Service {
	srvc := &Service{
		port: port,
		Name: name,
	}

	srvc.Reset()
	return srvc
}

func (service *Service) Reset() {
	service.packetHandlers = make(map[uint32]PacketHandler)
	service.peers = make(map[*protocol.CNPeer]interface{})
	service.started = make(chan struct{})
}

// may not be called while the service is running (eg. srvc.Start() has been called)
func (service *Service) AddPacketHandler(pktID uint32, handler PacketHandler) {
	service.packetHandlers[pktID] = handler
}

func (service *Service) Start() error {
	service.stop = make(chan struct{})
	service.stopped = make(chan struct{})
	peerConnections := make(chan chan *protocol.Event)
	go service.handleEvents(peerConnections)

	// open listener socket
	var err error
	service.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", service.port))
	if err != nil {
		return err
	}

	close(service.started) // signal that the service has started
	log.Printf("%s service hosted on %s:%d\n", service.Name, config.GetAnnounceIP(), service.port)
	for {
		conn, err := service.listener.Accept()
		if err != nil {
			fmt.Println(err)
			// we expect this to happen when the service is stopped
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}

		// create a new peer and pass it to the event loop
		eRecv := make(chan *protocol.Event)
		peer := protocol.NewCNPeer(conn)
		log.Printf("New peer %p connected to %s\n", peer, service.Name)
		peerConnections <- eRecv
		go peer.Handler(eRecv)
	}
}

// returns a channel that is closed when the service has started.
// this is useful if you need to do something after the service has started.
func (service *Service) Started() <-chan struct{} {
	return service.started
}

// returns a channel that is closed when the service has stopped.
// this is useful if you need to wait until the service has completely stopped.
func (service *Service) Stopped() <-chan struct{} {
	return service.stopped
}

// stops the service and disconnects all peers. OnDisconnect will be called
// for each peer.
func (service *Service) Stop() {
	close(service.stop)

	service.listener.Close()
}

// returns the stored uData for the peer.
// if the peer does not exist, nil is returned.
// NOTE: the peer map is not locked while accessing, if you're calling this
// outside of the service's event loop, you'll need to lock the peer map yourself.
func (service *Service) GetPeerData(peer *protocol.CNPeer) interface{} {
	return service.peers[peer]
}

// sets the stored uData for the peer.
// NOTE: the peer map is not locked while accessing, if you're calling this
// outside of the service's event loop, you'll need to lock the peer map yourself.
func (service *Service) SetPeerData(peer *protocol.CNPeer, uData interface{}) {
	service.peers[peer] = uData
}

// calls f for each peer in the service passing the peer and the stored uData.
// if f returns false, the iteration is stopped.
// NOTE: the peer map is not locked while iterating, if you're calling this
// outside of the service's event loop, you'll need to lock the peer map yourself.
func (service *Service) RangePeers(f func(peer *protocol.CNPeer, uData interface{}) bool) {
	for peer, uData := range service.peers {
		if !f(peer, uData) {
			break
		}
	}
}

// locks the peer map.
func (service *Service) Lock() {
	service.stateLock.Lock()
}

// unlocks the peer map.
func (service *Service) Unlock() {
	service.stateLock.Unlock()
}

// handleEvents is the main event loop for the service.
// it handles all events from the peers and calls the appropriate handlers.
func (service *Service) handleEvents(eRecv <-chan chan *protocol.Event) {
	poll := make([]reflect.SelectCase, 0, 4)

	// add the stop channel and the peer connection channel to our poll queue
	poll = append(poll, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(service.stop),
	})
	poll = append(poll, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(eRecv),
	})

	for {
		chosen, value, _ := reflect.Select(poll)
		if chosen == 0 {
			// stop

			// OnDisconnect handler might need to do something important
			service.Lock()
			service.RangePeers(func(peer *protocol.CNPeer, uData interface{}) bool {
				peer.Kill()
				service.disconnect(peer)
				return true
			})
			service.Unlock()

			// signal we have stopped
			close(service.stopped)
			return
		} else if chosen == 1 {
			// new peer, add it to our poll queue
			poll = append(poll, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(value.Interface()),
			})
		} else {
			// peer event
			event, ok := value.Interface().(*protocol.Event)
			if !ok {
				panic("invalid event type")
			}

			service.Lock()
			switch event.Type {
			case protocol.EVENT_CLIENT_DISCONNECT:
				// strip the peer from our poll queue
				poll = append(poll[:chosen], poll[chosen+1:]...)
				service.disconnect(value.Interface().(*protocol.Event).Peer)
			case protocol.EVENT_CLIENT_CONNECT:
				service.connect(event.Peer)
			case protocol.EVENT_CLIENT_PACKET:
				if err := service.handlePacket(event.Peer, event.PktID, protocol.NewPacket(event.Pkt)); err != nil {
					log.Printf("Error handling packet: %v", err)
					event.Peer.Kill()
				}

				// the packet buffer is given to us by the event, so we'll need to make sure to return it to the pool
				protocol.PutBuffer(event.Pkt)
			}
			service.Unlock()
		}
	}
}

func (service *Service) handlePacket(peer *protocol.CNPeer, typeID uint32, pkt protocol.Packet) error {
	uData := service.peers[peer]
	if hndlr, ok := service.packetHandlers[typeID]; ok {
		// fmt.Printf("Handling packet %x\n", typeID)
		if err := hndlr(peer, uData, pkt); err != nil {
			return err
		}
	} else {
		log.Printf("[WARN] unknown packet ID: %x\n", typeID)
	}

	return nil
}

func (service *Service) disconnect(peer *protocol.CNPeer) {
	if service.OnDisconnect != nil {
		uData := service.peers[peer]
		service.OnDisconnect(peer, uData)
	}

	log.Printf("Peer %p disconnected from %s\n", peer, service.Name)
	delete(service.peers, peer)
}

func (service *Service) connect(peer *protocol.CNPeer) {
	// default uData to nil, but if the service has an OnConnect
	// handler, use the result from that
	uData := interface{}(nil)
	if service.OnConnect != nil {
		uData = service.OnConnect(peer)
	}

	log.Printf("New peer %p connected to %s\n", peer, service.Name)
	service.SetPeerData(peer, uData)
}
