package waitgroup

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitGroup_Wrap(t *testing.T) {
	wg := New()

	wg.Wrap(func() {
		fmt.Println("before", "hahahaha!!!balabala...")
		time.Sleep(time.Second * 1)
		fmt.Println("after", "hahahaha!!!balabala...")
	})

	// block until wrap(func()) finish
	wg.Wait()
}
