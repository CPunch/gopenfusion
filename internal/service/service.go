package service

import (
	"fmt"
	"log"
	"net"
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
	eRecv          chan *protocol.Event
	packetHandlers map[uint32]PacketHandler
	peers          *sync.Map

	// OnDisconnect is called when a peer disconnects from the service.
	// uData is the stored value of the key/value pair in the peer map.
	// It may not be set while the service is running. (eg. srvc.Start() has been called)
	OnDisconnect func(peer *protocol.CNPeer, uData interface{})

	// OnConnect is called when a peer connects to the service.
	// return value is used as the value in the peer map.
	// It may not be set while the service is running. (eg. srvc.Start() has been called)
	OnConnect func(peer *protocol.CNPeer) (uData interface{})
}

func NewService(name string, port int) (*Service, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	service := &Service{
		listener:       listener,
		port:           port,
		Name:           name,
		eRecv:          make(chan *protocol.Event),
		packetHandlers: make(map[uint32]PacketHandler),
		peers:          &sync.Map{},
	}

	return service, nil
}

// may not be called while the service is running (eg. srvc.Start() has been called)
func (service *Service) AddPacketHandler(pktID uint32, handler PacketHandler) {
	service.packetHandlers[pktID] = handler
}

func (service *Service) Start() {
	log.Printf("%s service hosted on %s:%d\n", service.Name, config.GetAnnounceIP(), service.port)

	go service.handleEvents()
	for {
		conn, err := service.listener.Accept()
		if err != nil {
			log.Println("Connection error: ", err)
			return
		}

		peer := protocol.NewCNPeer(service.eRecv, conn)
		service.connect(peer)
	}
}

func (service *Service) handleEvents() {
	for event := range service.eRecv {
		switch event.Type {
		case protocol.EVENT_CLIENT_DISCONNECT:
			service.disconnect(event.Peer)
		case protocol.EVENT_CLIENT_PACKET:
			if err := service.handlePacket(event.Peer, event.PktID, protocol.NewPacket(event.Pkt)); err != nil {
				log.Printf("Error handling packet: %v", err)
				event.Peer.Kill()
			}

			// the packet buffer is given to us by the event, so we'll need to make sure to return it to the pool
			protocol.PutBuffer(event.Pkt)
		}
	}
}

func (service *Service) handlePacket(peer *protocol.CNPeer, typeID uint32, pkt protocol.Packet) error {
	uData, _ := service.peers.Load(peer)
	if hndlr, ok := service.packetHandlers[typeID]; ok {
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
		uData, _ := service.peers.Load(peer)
		service.OnDisconnect(peer, uData)
	}

	log.Printf("Peer %p disconnected from %s\n", peer, service.Name)
	service.peers.Delete(peer)
}

func (service *Service) connect(peer *protocol.CNPeer) {
	// default uData to nil, but if the service has an OnConnect
	// handler, use the result from that
	uData := interface{}(nil)
	if service.OnConnect != nil {
		uData = service.OnConnect(peer)
	}

	log.Printf("New peer %p connected to %s\n", peer, service.Name)
	service.peers.Store(peer, uData)
	go peer.Handler()
}

func (service *Service) SetPeerData(peer *protocol.CNPeer, uData interface{}) {
	service.peers.Store(peer, uData)
}

func (service *Service) RangePeers(f func(peer *protocol.CNPeer, uData interface{}) bool) {
	service.peers.Range(func(key, value any) bool {
		return f(key.(*protocol.CNPeer), value)
	})
}
