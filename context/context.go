package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func WithCancel() error {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker")
			case <-ctx.Done(): // 协程里面监听context的done channel, 如果Done, return出去，结束协程
				fmt.Println("roundtine exiting...", ctx.Err())
				return
			}
		}
	}(ctx)

	// 发出结束go协程的信号
	cancel()

	// 等待go协程执行完
	wg.Wait()

	return nil
}

func WithDeadline() {
	var (
		wg sync.WaitGroup
	)
	// 到期后执行cancel
	c, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c.Done():
				fmt.Println(c.Deadline())
				return
			}
		}
	}()

	//cancel()

	wg.Wait()
}

func WithTimeout() {
	var (
		wg sync.WaitGroup
	)

	// 超时后执行cancel
	c, _ := context.WithTimeout(context.Background(), time.Second*2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c.Done():
				fmt.Println(c.Deadline())
				return
			}
		}
	}()

	// cancel可以主动执行
	// cancel()

	wg.Wait()
}

func WithValue() {
	var (
		wg sync.WaitGroup
	)

	ctx, _ := context.WithCancel(context.Background())
	c := context.WithValue(ctx, "name", "derek")

	// 获取key值，用于标识context
	fmt.Println(c.Value("name"))

	wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker")
			case <-ctx.Done(): // 协程里面监听context的done channel, 如果Done, return出去，结束协程
				fmt.Println("roundtine exiting...", ctx.Err())
				return
			}
		}
	}(ctx)

	// 结束go协程
	cancel()

	wg.Wait()
}
