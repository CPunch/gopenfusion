package testutil_test

import (
	"sync"
	"testing"
	"time"

	"github.com/CPunch/gopenfusion/internal/testutil"
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
	is.True(!testutil.WaitWithTimeout(wg, 500*time.Millisecond)) // timeout should occur
	is.True(testutil.WaitWithTimeout(wg, 750*time.Millisecond))  // timeout shouldn't occur
}

func TestSelectWithTimeout(t *testing.T) {
	is := is.New(t)
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	is.True(!testutil.SelectWithTimeout(ch, 500*time.Millisecond)) // timeout should occur
	is.True(testutil.SelectWithTimeout(ch, 750*time.Millisecond))  // timeout shouldn't occur
}
