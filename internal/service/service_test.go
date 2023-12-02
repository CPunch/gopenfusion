package service_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/CPunch/gopenfusion/cnet"
	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/CPunch/gopenfusion/internal/service"
	"github.com/matryer/is"
)

var (
	srvcPort int
)

const (
	timeout       = 2
	maxDummyPeers = 5
)

func selectWithTimeout(ch <-chan struct{}, seconds int) bool {
	select {
	case <-ch:
		return true
	case <-time.After(time.Duration(seconds) * time.Second):
		return false
	}
}

func waitWithTimeout(wg *sync.WaitGroup, seconds int) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	return selectWithTimeout(done, seconds)
}

func TestMain(m *testing.M) {
	var err error
	srvcPort, err = service.RandomPort()
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestService(t *testing.T) {
	is := is.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	srvc := service.NewService(ctx, "TEST", srvcPort)
	wg := sync.WaitGroup{}

	// shutdown service when test is done
	defer func() {
		cancel()
		is.True(selectWithTimeout(srvc.Stopped(), timeout)) // wait for service to stop with timeout
	}()

	// our dummy packet handler
	wg.Add(maxDummyPeers)
	srvc.AddPacketHandler(0x1234, func(peer *cnet.Peer, pkt protocol.Packet) error {
		log.Printf("Received packet %#v", pkt)
		wg.Done()
		return nil
	})

	// wait for all dummy peers to connect and disconnect
	wg.Add(maxDummyPeers)
	srvc.OnConnect = func(peer *cnet.Peer) {
		wg.Done()
	}

	wg.Add(maxDummyPeers)
	srvc.OnDisconnect = func(peer *cnet.Peer) {
		wg.Done()
	}

	// run service
	go func() { is.NoErr(srvc.Start()) }()              // srvc.Start error
	is.True(selectWithTimeout(srvc.Started(), timeout)) // wait for service to start with timeout

	wg.Add(maxDummyPeers * 2) // 2 wg.Done() per peer for receiving packets
	for i := 0; i < maxDummyPeers; i++ {
		go func() {
			// make dummy client
			conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", srvcPort))
			is.NoErr(err) // net.Dial error

			peer := cnet.NewPeer(ctx, conn)
			go func() {
				defer peer.Kill()

				// send dummy packets
				for i := 0; i < 2; i++ {
					is.NoErr(peer.Send(0x1234)) // peer.Send error
				}
			}()

			// we wait until Handler gracefully exits (peer was killed)
			peer.Handler(make(chan *cnet.PacketEvent))
			wg.Done()
		}()
	}

	is.True(waitWithTimeout(&wg, timeout)) // wait for all dummy peers to be done with timeout
}
