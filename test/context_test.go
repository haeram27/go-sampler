package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestValueContext(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "key1", "value1")
	paretCtx := context.WithValue(baseCtx, "key2", "value2")
	childCtx := context.WithValue(paretCtx, "key3", "value3")

	t.Log(childCtx.Value("key1")) // value1
	t.Log(childCtx.Value("key2")) // value2
	t.Log(childCtx.Value("key3")) // value3
	t.Log(baseCtx.Value("key2"))  // <nil>
	t.Log(paretCtx.Value("key3")) // <nil>
}

func BlockingFnA(ctx context.Context, printCh <-chan int, funcName string) {
	for {
		select {
		case <-ctx.Done(): // receive context Done channel by invoking CancelFunc()
			fmt.Printf("%s :: (P)%s: ctx.Done() is called\n", funcName, ctx.Value("parentFn"))
			if err := ctx.Err(); err != nil {
				fmt.Printf("%s :: err: %s\n", funcName, err)
			}
			fmt.Printf("%s :: context is done. finished.\n", funcName)
			return // terminate go-routine

		case num := <-printCh:
			fmt.Printf("%s :: %d\n", funcName, num)

		default: // check select in every second than nano time in case there is no channel event
			time.Sleep(1 * time.Second)
		}
	}
}

func TestCancelContext(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestCancelContext")
	ctx, cancelCtx := context.WithCancel(baseCtx)

	printCh := make(chan int)

	go BlockingFnA(ctx, printCh, "BlockingFn1") // run go-routine

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx() // send data into ctx.Done() channel immediately
	time.Sleep(4 * time.Second)

	fmt.Println("TestCancelContext:: finished")
}

func TestCancelContextM(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestCancelContext")
	ctx, cancelCtx := context.WithCancel(baseCtx)

	printCh := make(chan int)

	go BlockingFnA(ctx, printCh, "BlockingFn1") // run go-routine
	go BlockingFnA(ctx, printCh, "BlockingFn2") // run go-routine
	go BlockingFnA(ctx, printCh, "BlockingFn3") // run go-routine

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx() // send data into ctx.Done() channel immediately
	time.Sleep(4 * time.Second)

	fmt.Println("TestCancelContext:: finished")
}

func TestTimeoutContext(t *testing.T) {
	// baseCtx SHOULD be made per any sending/receiving HTTP Request
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestTimeoutContext")

	// 1500 msec, same as ctx.WithDeadline(ctx, time.Now().Add(1500 * time.Millisecond))
	ctx, cancelCtx := context.WithTimeout(baseCtx, 1500*time.Millisecond)
	defer cancelCtx()

	printCh := make(chan int)

	go BlockingFnA(ctx, printCh, "BlockingFn1") // go-routine

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Println("TestTimeoutContext::  ctx.Done() is called")
			break
		}
	}

	fmt.Println("TestTimeoutContext:: finished")
}

func TestDeadlineContext(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestDeadlineContext")

	// 1500 msec, same as ctx.WithTimeout(ctx, 1500 * time.Millisecond)
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(baseCtx, deadline)
	defer cancelCtx()

	printCh := make(chan int)

	go BlockingFnA(ctx, printCh, "TestDeadlineContext") // go-routine

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Println("TestDealineContext::  ctx.Done() is called")
			break
		}

	}

	fmt.Println("TestDealineContext:: finished")
}

func BlockingFnD(parent context.Context, child context.Context) {
	for {
		select {
		case <-parent.Done(): // receive context Done channel by invoking CancelFunc()
			fmt.Printf("BlockingFnD:: (P)%s: parent.Done() is called\n", parent.Value("parentFn"))
			if err := parent.Err(); err != nil {
				fmt.Printf("BlockingFnD:: parent err: %s\n", err)
			}
			time.Sleep(1 * time.Second)

		case <-child.Done(): // receive context Done channel by invoking CancelFunc()
			fmt.Printf("BlockingFnD:: (P)%s: child.Done() is called\n", child.Value("parentFn"))
			if err := child.Err(); err != nil {
				fmt.Printf("BlockingFnD:: child err: %s\n", err)
			}
			time.Sleep(1 * time.Second)

		default: // check select in every second than nano time in case there is no channel event
			time.Sleep(1 * time.Second)
		}
	}
}

func TestCancelContextCancelParent(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestCancelContextD")
	parent, cancelParent := context.WithCancel(baseCtx)
	defer cancelParent()
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()

	go BlockingFnD(parent, child) // go-routine

	/* cancelParent() make event of child.Done() channel */
	cancelParent() // send data into ctx.Done() channel immediately
	time.Sleep(4 * time.Second)

	fmt.Println("TestCancelContextD:: finished")
}

func TestCancelContextCancelChild(t *testing.T) {
	baseCtx := context.WithValue(context.Background(), "parentFn", "TestCancelContextD")
	parent, cancelParent := context.WithCancel(baseCtx)
	defer cancelParent()
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()

	go BlockingFnD(parent, child) // go-routine

	/* cancelChild() CANNOT make event of parent.Done() channel */
	cancelChild() // send data into ctx.Done() channel immediately
	time.Sleep(4 * time.Second)

	fmt.Println("TestCancelContextD:: finished")
}
