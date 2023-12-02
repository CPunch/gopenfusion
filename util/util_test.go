package util_test

import (
	"sync"
	"testing"
	"time"

	"github.com/CPunch/gopenfusion/util"
	"github.com/matryer/is"
)

func TestWaitWithTimeout(t *testing.T) {
	is := is.New(t)
	wg := &sync.WaitGroup{}
	go func() {
		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	wg.Add(1)
	is.True(!util.WaitWithTimeout(wg, 500*time.Millisecond)) // timeout should occur
	is.True(util.WaitWithTimeout(wg, 750*time.Millisecond))  // timeout shouldn't occur
}

func TestSelectWithTimeout(t *testing.T) {
	is := is.New(t)
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	is.True(!util.SelectWithTimeout(ch, 500*time.Millisecond)) // timeout should occur
	is.True(util.SelectWithTimeout(ch, 750*time.Millisecond))  // timeout shouldn't occur
}
