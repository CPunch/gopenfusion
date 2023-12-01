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

	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/CPunch/gopenfusion/internal/service"
)

var (
	srvcPort int
)

const (
	timeout       = 2
	maxDummyPeers = 5
)

func waitWithTimeout(wg *sync.WaitGroup, seconds int) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
		return true
	case <-time.After(time.Duration(seconds) * time.Second):
		return false
	}
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
	ctx, cancel := context.WithCancel(context.Background())
	srvc := service.NewService(ctx, "TEST", srvcPort)

	// waitgroup to wait for test packet handler to be called
	wg := sync.WaitGroup{}

	srvc.AddPacketHandler(0x1234, func(peer *protocol.CNPeer, pkt protocol.Packet) error {
		log.Printf("Received packet %#v", pkt)
		wg.Done()
		return nil
	})

	go func() {
		if err := srvc.Start(); err != nil {
			t.Error(err)
		}
	}()

	// wait for service to start
	<-srvc.Started()
	wg.Add(maxDummyPeers * 3) // 2 wg.Done() calls per dummy peer
	for i := 0; i < maxDummyPeers; i++ {
		go func() {
			// make dummy client
			conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", srvcPort))
			if err != nil {
				t.Error(err)
			}

			peer := protocol.NewCNPeer(ctx, conn)
			go func() {
				defer peer.Kill()

				// send dummy packets
				for i := 0; i < 2; i++ {
					if err := peer.Send(0x1234); err != nil {
						t.Error(err)
					}
				}
			}()

			// we wait until Handler gracefully exits (peer was killed)
			peer.Handler(make(chan *protocol.PacketEvent))
			wg.Done()
		}()
	}

	if !waitWithTimeout(&wg, timeout) {
		t.Error("failed to wait for packet handler to be called")
	}

	cancel()
	<-srvc.Stopped()
}
