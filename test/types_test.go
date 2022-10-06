package test

import (
	"testing"
)

func TestDefautValueOfTyep(t *testing.T) {

	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f uint
		g uint8
		h uint16
		i uint32
		j uint64
		k uintptr
		l byte // byte is alias for uint8
		m rune // rune is alias for int32, represents a Unicode code point
		n complex64
		o complex128
		p bool // false
		q *int // nil
		r interface{}
		s struct{} // {}
		u []int
		v map[string]interface{}
		w chan int
		//y func(){}
		//? interface{}{} // interface{}{}??
		//? interface // interface??
	)

	t.Log("int: ", a)                    // 0
	t.Log("int8: ", b)                   // 0
	t.Log("int16: ", c)                  // 0
	t.Log("int32: ", d)                  // 0
	t.Log("int64: ", e)                  // 0
	t.Log("uint: ", f)                   // 0
	t.Log("uint8: ", g)                  // 0
	t.Log("uint16: ", h)                 // 0
	t.Log("uint32: ", i)                 // 0
	t.Log("uint64: ", j)                 // 0
	t.Log("uintptr: ", k)                // 0
	t.Log("byte: ", l)                   // 0
	t.Log("rune: ", m)                   // 0
	t.Log("complex64: ", n)              // (0+0i)
	t.Log("complex128: ", o)             // (0+0i)
	t.Log("bool: ", p)                   // false
	t.Log("pointer: ", q)                // <nil>
	t.Log("interface{}: ", r)            // <nil>
	t.Log("struct{}: ", s)               // {}
	t.Log("[]int{}: ", u)                // []
	t.Log("map[string]interface{}: ", v) // {}
	t.Log("chan: ", w)                   // map[]
}

func TestNilType(t *testing.T) {
	// nil can be pointer, slice, map, channel, func, empty interface
	// https://go101.org/article/nil.html

	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// These lines are equivalent to the above lines.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// Retrieving Elements From Nil Maps Will Not Panic
	t.Log((map[string]int)(nil)["key"]) // 0
	t.Log((map[int]bool)(nil)[123])     // false
	t.Log((map[int]*int64)(nil)[123])   //

	//*new(T) Results a Nil T Value if the Zero Value of Type T Is Represented With the Predeclared nil Identifier
	t.Log(*new(*int) == nil)         // true
	t.Log(*new([]int) == nil)        // true
	t.Log(*new(map[int]bool) == nil) // true
	t.Log(*new(chan string) == nil)  // true
	t.Log(*new(func()) == nil)       // true
	t.Log(*new(interface{}) == nil)  // true
}
