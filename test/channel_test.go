package test

import (
	"fmt"
	"testing"
	"time"
)

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
