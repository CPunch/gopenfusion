package service_test

import (
	"fmt"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/CPunch/gopenfusion/internal/service"
)

var (
	srvc     *service.Service
	srvcPort int
)

const (
	timeout       = 5
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

	srvc = service.NewService("TEST", srvcPort)
	os.Exit(m.Run())
}

func TestService(t *testing.T) {
	// waitgroup to wait for test packet handler to be called
	wg := sync.WaitGroup{}

	srvc.AddPacketHandler(0x1234, func(peer *protocol.CNPeer, uData interface{}, pkt protocol.Packet) error {
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
	wg.Add(maxDummyPeers)
	for i := 0; i < maxDummyPeers; i++ {
		go func() {
			// make dummy client
			conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", srvcPort))
			if err != nil {
				t.Error(err)
			}

			peer := protocol.NewCNPeer(conn)
			defer peer.Kill()
			// send dummy packet
			if err := peer.Send(0x1234); err != nil {
				t.Error(err)
			}
		}()
	}

	if !waitWithTimeout(&wg, timeout) {
		t.Error("timeout waiting for packet handler to be called")
	}
	srvc.Stop()
	<-srvc.Stopped()
}
