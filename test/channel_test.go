package test

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestChannelSize(t *testing.T) {
	ch := make(chan int)
	t.Log(cap(ch)) // 0
	t.Log(len(ch)) // 0

	go func() {
		<-time.After(3)
		t.Log("cap: ", cap(ch)) // 0
		t.Log("len: ", len(ch)) // 0
		t.Log("pop: ", <-ch)    // 1  , release blocking in main goroutine
	}()
	t.Log("push  1")
	ch <- 1 // blocking main goroutine

	t.Log("finish")
}

func TestChannelPop(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 2; i++ {
			ch <- i
			t.Log("send: ", i)
		}
		close(ch)
	}()

	for cnt := 0; cnt < 10; cnt++ {
		v, ok := <-ch
		if ok != false {
			t.Log("recv: ", v, ", ", ok)
		}
	}

	t.Log("finish")
}

func TestChannelPopForRange(t *testing.T) {
	ch1 := make(chan int, 100)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 140; i++ {
			ch1 <- i
			t.Log("send: ", i)
		}
		close(ch1)
	}()

	for v := range ch1 {
		t.Log("recv: ", v)
	}

	t.Log("finish")
}

func TestChannelForSelect(t *testing.T) {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		wg.Done()
	}()
	wg.Wait()

Loop:
	for {
		select {
		case val := <-ch1:
			t.Log(val)
			if val == 3 {
				break Loop
			}
		case <-ch2:
		}
	}

	t.Log("finish")
}

func TestChannelRangeAndHidenSize(t *testing.T) {

	// channel can be store cap+1 size of date
	ch := make(chan int, 3) // make channel with capacity 3 and hiden 1
	t.Log("len: ", len(ch)) // 0
	t.Log("cap: ", cap(ch)) // 3

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Second)
		t.Log("len: ", len(ch)) // 0
		for i := range ch {     // range will pop elements in chan at a time
			t.Log("rcv: ", i) // rcv 1-4 at a time but cap of ch was 3
		}
		wg.Done()
	}()

	t.Log("1")
	ch <- 1
	t.Log("2")
	ch <- 2
	t.Log("3")
	ch <- 3
	t.Log("4")
	ch <- 4 // block after push 4
	t.Log("5")
	ch <- 5
	t.Log("6")
	ch <- 6
	t.Log("====")

	close(ch)
	wg.Wait()
}

func request(s string) string {
	time.Sleep(time.Duration(rand.Int()%5) * time.Second)
	return s
}

func TestChannelWithGoroutine(t *testing.T) {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	go func() { responses <- request("africas.gopl.io") }()
	t.Log(<-responses) // 가장 빠른 응답 반환

	fmt.Print("TestChannel() finished\n")
}

func TestChannelRaceInSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch1 <- 1:
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch2 <- 1:
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var ch1sum, ch2sum int
		for i := 0; i < 1000000; i++ {
			select {
			case <-ch1:
				ch1sum++
			case <-ch2:
				ch2sum++
			}
		}
		t.Logf("%d %d\n", ch1sum, ch2sum)
		wg.Done()
	}()
	wg.Wait()
	cancel()
	t.Log("TestChannelRaceInSelect() finished\n")
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 1
	close(ch)
	go func() {
		t.Log("start go\n")
		for v := range ch {
			t.Logf("print: %v", v)
		}
	}()

	time.Sleep(10 * time.Second)

	fmt.Print("TestChannelRaceInSelect() finished\n")
}
