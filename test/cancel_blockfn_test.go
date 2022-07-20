package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func BlockingFnB(ctx context.Context, startCh <-chan bool, caller string) {
	for {
		select {
		case <-ctx.Done(): // receive context Done channel by invoking CancelFunc()
			fmt.Printf("%s :: (P)%s: ctx.Done() is called\n", caller, ctx.Value("parentFn"))
			if err := ctx.Err(); err != nil {
				fmt.Printf("%s :: err: %s\n", caller, err)
			}
			fmt.Printf("%s :: finished\n", caller)
			return // terminate go-routine
		case <-startCh:
			for {
				// Do something blocking here
				fmt.Printf("%s is blocking...\n", caller)
				time.Sleep(1 * time.Second)
			}
		default: // check select in every second than nano time in case there is no channel event
			time.Sleep(1 * time.Second)
		}
	}
}

func TestCancelBlockingFn(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestCancelBlockingFn")
	ctx, cancelCtx := context.WithCancel(baseCtx)
	startCh := make(chan bool)
	go BlockingFnB(ctx, startCh, "BlockingFn1") // run go-routine
	startCh <- true

	time.Sleep(3 * time.Second)
	cancelCtx() // send data into ctx.Done() channel immediately
	time.Sleep(3 * time.Second)
	fmt.Println("TestCancelBlockingFn:: finished")
}
