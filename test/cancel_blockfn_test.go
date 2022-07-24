package test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func DoSomething1(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done(): // check cancel event!! receive context Done channel by invoking CancelFunc()!!!
			fmt.Printf("ctx.Done() is called\n")
			if err := ctx.Err(); err != nil {
				fmt.Printf("err: %s\n", err)
				return err // terminate go-routine
			} else {
				return errors.New("canceled") // terminate go-routine
			}

		default: // default makes let runtime go out of select scope
			fmt.Println("default")
		}

		// Do Something here
		fmt.Println("out of select")
	}
}

func DoSomething2(ctx context.Context) error {
	// Do Something here
	fmt.Println("out of select")

	select {
	case <-ctx.Done(): // check cancel event!! receive context Done channel by invoking CancelFunc()!!!
		fmt.Printf("ctx.Done() is called\n")
		if err := ctx.Err(); err != nil {
			fmt.Printf("err: %s\n", err)
			return err // terminate go-routine
		} else {
			return errors.New("canceled") // terminate go-routine
		}

	default: // default makes let runtime go out of select scope
		fmt.Println("default")
	}

	return nil
}

func TestDoSomething1(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestDoSomething1")
	ctx, cancelCtx := context.WithCancel(baseCtx)

	go DoSomething1(ctx) // run go-routine

	time.Sleep(10 * time.Second)
	cancelCtx() // send data into ctx.Done() channel immediately

	fmt.Println("TestCancelBlockingFn:: finished")
}
