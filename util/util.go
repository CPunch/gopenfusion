package util

import (
	"sync"
	"time"
)

func GetTime() uint64 {
	return uint64(time.Now().UnixMilli())
}

func SelectWithTimeout(ch <-chan struct{}, seconds int) bool {
	select {
	case <-ch:
		return true
	case <-time.After(time.Duration(seconds) * time.Second):
		return false
	}
}

func WaitWithTimeout(wg *sync.WaitGroup, seconds int) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	return SelectWithTimeout(done, seconds)
}
