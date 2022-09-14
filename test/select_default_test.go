package test

import (
	"fmt"
	"testing"
)

func g1(ch chan int) {
	ch <- 42
}

func TestSelectDefault(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go g1(ch1)
	go g1(ch2)

	/*
		this will make time to get event from channel in select state
		If uncomment time.Sleep() invoking, the result is changed.
	*/
	// time.Sleep(1 * time.Second)

	select {
	case v1 := <-ch1:
		fmt.Println("Got: ", v1)
	case v2 := <-ch2:
		fmt.Println("Got: ", v2)

	/*
		if default block is defined in select block then select block will never block goroutine.
		if whole above case expression is blocking function
		then default statement is called because there is no arrived event in above cases.
		and select statement block is escaped after running default statement block.
	*/
	default:
		fmt.Println("The default case!")
	}

}

func g2(ch chan int) {
	for {
		ch <- 43
	}
}

func TestSelectFor(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go g2(ch1)
	go g2(ch2)

	cnt1 := 0
	cnt2 := 0

	for i := 0; i < 1000000; i++ {
		select {
		case <-ch1:
			cnt1++
		case <-ch2:
			cnt2++
		}
	}
	t.Log(cnt1, cnt2) // 501280 498720
}

func TestSelectRace(t *testing.T) {
	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)

	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			if i%2 == 0 {
				ch2 <- i
			} else {
				ch1 <- i
			}
		}
		// wg.Done()
	}()
	// wg.Wait()

LOOP:
	for {
		select {
		case v := <-ch1:
			t.Log("ch1: ", v)
			if v == 999 {
				break LOOP
			}
		case v := <-ch2:
			t.Log("ch2: ", v)
		default:
			t.Log("default")
		}
	}

	t.Log("finish")
}
