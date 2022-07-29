package test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func DoRepeatedJob(ctx context.Context) error {
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

		// Do repeated work here
		fmt.Println("repeat !!!")
	}
}

func TestDoRepeatedJob(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	go DoRepeatedJob(ctx) // run go-routine

	time.Sleep(10 * time.Second)
	cancelCtx() // send data into ctx.Done() channel immediately

	fmt.Println("TestDoRepeatedJob:: finished")
}

func DoBlockingJob(ctx context.Context) error {
	// Do blocking job here
	for i := 0; i < 5; i++ {
		fmt.Println("blocking !!!")
		time.Sleep(1 * time.Second)
	}

	select {
	case <-ctx.Done(): // check cancel event!! receive context Done channel by invoking CancelFunc()!!!
		fmt.Printf("ctx.Done() is called\n")
		if err := ctx.Err(); err != nil {
			fmt.Printf("context is canceled. err: %s\n", err)
			return err // terminate go-routine
		} else {
			return errors.New("context is canceled") // terminate go-routine
		}

	default: // default makes let runtime go out of select scope
		fmt.Println("default")
	}

	return nil
}

func TestDoBlockingJob(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	go DoBlockingJob(ctx) // run go-routine

	time.Sleep(10 * time.Second)
	cancelCtx() // send data into ctx.Done() channel immediately

	fmt.Println("TestDoBlockingJob:: finished")
}
