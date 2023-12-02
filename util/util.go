package util

import (
	"sync"
	"time"
)

func GetTime() uint64 {
	return uint64(time.Now().UnixMilli())
}

func SelectWithTimeout(ch <-chan struct{}, timeout time.Duration) bool {
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

func WaitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	return SelectWithTimeout(done, timeout)
}
