package test

import (
	"reflect"
	"testing"
	"unsafe"
)

type STRUCT struct {
	i int64
	j rune
}

type UNEFFICIENT struct {
	a int8  // 1byte, making 7byte padding
	b int64 // 8byte
	c int8  // 1byte, making 7byte padding
	d int64 // 8byte
}

func TestUnsafePointer(t *testing.T) {
	var a int64
	var strt STRUCT

	var uintptrT1 uintptr
	// uintptrT1 = uintptr(&a)              // INVALID: cannot convert &a (value of type *int) to uintptr
	uintptrT1 = uintptr(unsafe.Pointer(&a)) // OK
	t.Logf("%p", &a)                        // 0xc0000182f8
	t.Log(unsafe.Pointer(&a))               // 0xc0000182f8
	t.Log(uintptrT1)                        // 824633819896

	t.Log(reflect.TypeOf(unsafe.Pointer(&a)).Kind())          // unsafe.Pointer, HEX expr of ponter address
	t.Log(reflect.TypeOf(uintptr(unsafe.Pointer(&a))).Kind()) // uintptr, DEC(uint) expr of ponter address
	t.Log(unsafe.Pointer(&a))                                 // 0xc0000182d8, HEX expr of ponter address of a
	t.Log(uintptr(unsafe.Pointer(&a)))                        // 824633819864, DEC(uint) expr of ponter address of a

	t.Log(unsafe.Pointer(&strt))          // 0xc0000182e0, address(uintptr(HEX)) of strt, WARINING: &strt == &strt.i
	t.Log(uintptr(unsafe.Pointer(&strt))) // 824633819872, address(uintptr(DEC)) of strt

	// unsafe.Pointer(&strt.i) == unsafe.Pointer(&strt)
	t.Log(unsafe.Pointer(&strt.i))                                  // 0xc0000182e0, address(unsafe.Pointer(HEX)) of strt.i
	t.Log(uintptr(unsafe.Pointer(&strt)) + unsafe.Offsetof(strt.i)) // 824633819872, address(uintptr(DEC)) of strt.i
	t.Log(unsafe.Offsetof(strt.i))                                  // 0
}

func TestUnsafePointerCasting(t *testing.T) {
	var a int32
	var b *uint32

	b = (*uint32)(unsafe.Pointer(&a))

	t.Log(reflect.TypeOf(&a)) // *int32
	t.Log(reflect.TypeOf(b))  // *uint32
	t.Logf("%p", &a)          //0xc000018368
	t.Logf("%p", b)           //0xc000018368
}

func TestUnsafeOffsetof(t *testing.T) {
	var strt STRUCT

	t.Log(unsafe.Pointer(&strt))          // 0xc0000182e0, address(uintptr(HEX)) of strt, WARINING: &strt == &strt.i
	t.Log(uintptr(unsafe.Pointer(&strt))) // 824633819872, address(uintptr(DEC)) of strt

	// unsafe.Pointer(&strt.i) == unsafe.Pointer(&strt)
	t.Log(unsafe.Pointer(&strt.i))                                  // 0xc0000182e0, address(unsafe.Pointer(HEX)) of strt.i
	t.Log(uintptr(unsafe.Pointer(&strt)) + unsafe.Offsetof(strt.i)) // 824633819872, address(uintptr(DEC)) of strt.i
	t.Log(unsafe.Offsetof(strt.i))                                  // 0

	t.Log(unsafe.Pointer(&strt.j))                                  // 0xc0000182e8, address(unsafe.Pointer(HEX)) of strt.j
	t.Log(uintptr(unsafe.Pointer(&strt)) + unsafe.Offsetof(strt.j)) // 824633819880, address(uintptr(DEC)) of strt.j
	t.Log(unsafe.Offsetof(strt.j))                                  // 8 == distance between start of strt and start of strt.j == size of strt.i
	u := uintptr(unsafe.Pointer(&strt))
	offset := unsafe.Offsetof(strt.j)
	t.Log(u + offset) // 0xc0000182e8
}

func TestUnsafeSizeof(t *testing.T) {
	var a int64
	var strt STRUCT

	// func Sizeof(x ArbitraryType) uintptr
	// return size in bytes of type of variable x

	t.Log(unsafe.Sizeof(a))      // 8
	t.Log(unsafe.Sizeof(strt))   // 16
	t.Log(unsafe.Sizeof(strt.i)) // 8
	t.Log(unsafe.Sizeof(strt.j)) // 4
}

func TestUnsafeAlignof(t *testing.T) {
	var a int64
	var strt STRUCT
	// func Alignof(x ArbitraryType) uintptr
	// allignment? make minimum padding against struct instance
	t.Log(reflect.TypeOf(a).Align()) // 8
	t.Log(unsafe.Alignof(a))         // 8
	t.Log(unsafe.Alignof(strt))      // 4,  Warning : struct variable argument is calculated as first element of struct
	t.Log(unsafe.Alignof(strt.i))    // 4
	t.Log(unsafe.Alignof(strt.j))    // 4

	var un UNEFFICIENT
	t.Log(unsafe.Sizeof(un))  // 32
	t.Log(unsafe.Alignof(un)) // 8
}

func TestUnsafeSlice(t *testing.T) {
	s := make([]int, 10)
	s[0] = 10
	t.Log(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&s)).Data))
	t.Log(reflect.TypeOf(s))
	t.Log(cap(s))
	t.Log(len(s))

	// create new int sliece
	ss := unsafe.Slice(new(int), 10)
	t.Log(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&ss)).Data))
	t.Log(reflect.TypeOf(ss))
	t.Log(cap(ss))
	t.Log(len(ss))

	// convert int array to int sliece
	// Slice(ptr, len) is equivalent to (*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
	arr := [...]int{20, 21, 23}
	sss := unsafe.Slice(&arr[0], 20) // equivalent to sss := (*[20]int)(unsafe.Pointer(&arr[0]))[:]
	t.Logf("%p", &arr)
	t.Log(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&sss)).Data))
	t.Log(reflect.TypeOf(arr)) // [3]int
	t.Log(reflect.TypeOf(sss)) // []int
	t.Log(cap(sss))
	t.Log(len(sss))
}
