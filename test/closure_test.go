package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
In Go lang, basically closure is accessed as reference unless closure variable is NOT redeclared.
**/

func TestGoCapture(t *testing.T) {
	varOutter := "ref"                       // new instance
	fmt.Printf("original: %v\n", &varOutter) // original: 0xc00006b030

	var wg sync.WaitGroup
	print := func(name string) {
		fmt.Printf("%s: %v\n", name, &varOutter) // accessed as reference
		wg.Done()
	}

	wg.Add(2)
	go print("1") // 1: 0xc00006b030
	go print("2") // 2: 0xc00006b030
	wg.Wait()
}

func TestGoCapture2(t *testing.T) {
	varOutter := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			varOutter++                      // outter variable can be changed from closure
			fmt.Printf("in %v\n", varOutter) // accessed as reference
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		varOutter++
		fmt.Printf("out %v\n", varOutter)
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func TestClosureCaptureReDecl(t *testing.T) {
	redeclare := 1                                         // outter
	fmt.Printf("outter: %v,  %v\n", redeclare, &redeclare) // outter: 1,  0xc0000ba288

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		redeclare := 10                                       // inner is re-declared
		fmt.Printf("inner: %v,  %v\n", redeclare, &redeclare) // inner: 10,  0xc0000ba2a0
		wg.Done()
	}()
	wg.Wait()
}

func adder() func(int) {
	varOutter := 0 // new instance
	fmt.Printf("original: %v\n", &varOutter)
	return func(x int) {
		varOutter += x
		fmt.Printf("%d: %v\n", varOutter, &varOutter) // accessed as reference
	}
}

func TestClosureCapture(t *testing.T) {
	pos, neg := adder(), adder() // original: 0xc000128228, original: 0xc000128230
	pos(1)                       // 1: 0xc000128228
	neg(-2 * 1)                  // -2: 0xc000128230
}
