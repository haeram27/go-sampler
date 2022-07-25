package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestChannelSize(t *testing.T) {
	ch := make(chan int, 0)
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

func GoRotineRaceInSelect(i1 <-chan int, i2 <-chan int) {
	fmt.Printf("goroutine() started\n")
	for {
		select {
		case v := <-i1:
			fmt.Printf("[i1]: %d\n", v)
			time.Sleep(1 * time.Second)

		case v := <-i2:
			fmt.Printf("[i2]: %d\n", v)
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("goroutine() finished\n")
		return
	}
}

func TestChannelRaceInSelect(t *testing.T) {
	i1 := make(chan int, 1)
	i2 := make(chan int, 1)

	fmt.Print("Channel Input start\n")
	i1 <- 1
	i2 <- 2
	fmt.Print("Channel Input finished\n")

	go GoRotineRaceInSelect(i1, i2) // run go-routine

	fmt.Print("TestChannelRaceInSelect() finished\n")
}
