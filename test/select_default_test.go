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
		select doesn't have blocking function.
		default case is called when there is no arrived event in above cases.
	*/
	default:
		fmt.Println("The default case!")
	}

}

func g2(ch chan int) {
	for true {
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

	for i := 0; i < 40000; i++ {
		select {
		case <-ch1:
			cnt1++
		case <-ch2:
			cnt2++
		}
	}
	t.Log(cnt1, cnt2) // 20083 19917
}
