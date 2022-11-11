package test

import (
	"reflect"
	"testing"
	"unsafe"
)

type STRUCT struct { // 12 byte
	i int64   // 8 byte
	j rune    // 4 byte
	k *uint64 // pointer is system arch length
}

type STRUCT2 struct { // 12 byte
	i rune    // 4 byte
	j int64   // 8 byte
	k *uint32 // pointer is system arch length
}

type STRUCT3 struct { // 7 byte
	i uint8  // 1 byte
	j uint16 // 2 byte
	k uint32 // 4 byte
}

type UNEFFICIENT struct { // 18 byte
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

func TestUnsafePointerForceCasting(t *testing.T) {
	var a int32
	var b *uint32

	b = (*uint32)(unsafe.Pointer(&a))

	t.Log(reflect.TypeOf(&a)) // *int32
	t.Log(reflect.TypeOf(b))  // *uint32
	t.Logf("%p", &a)          // 0xc000018368
	t.Logf("%p", b)           // 0xc000018368
}

func TestDecapsulationByForceCasting(t *testing.T) {
	// convert slice pointer to reflect.SliceHeader pointer
	s := make([]int, 10)
	s[0] = 9
	s[1] = 8
	s[2] = 7
	t.Logf("s.type: %v", reflect.TypeOf(s))
	t.Logf("s.pointer: %p", &s)
	t.Logf("s[0].pointer: %p", &s[0])
	t.Logf("s.array: %v", s) // [9 8 7 0 0 0 0 0 0 0]
	t.Logf("len(s): %v", len(s))
	t.Logf("cap(s): %v", cap(s))

	srefhead := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	t.Logf("srefhead.pointer: %p", srefhead)                   // same as s.pointer
	t.Logf("srefhead.Data: %p", unsafe.Pointer(srefhead.Data)) // same as s[0].pointer and s.data
	arr := (*[10]int)(unsafe.Pointer(srefhead.Data))
	t.Logf("srefhead.array: %v", *arr)       // array contents in slice, [9 8 7 0 0 0 0 0 0 0]
	t.Logf("srefhead.Len: %v", srefhead.Len) // same as len(s)
	t.Logf("srefhead.Cap: %v", srefhead.Cap) // same as cao(s)
}

func TestUnsafeSlice(t *testing.T) {
	// unsafe.Slice() converts array to slice with new size(!!)
	// unsafe.Slice(ptr, len) is equivalent to (*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
	arr := [...]int{9, 8}
	ss := unsafe.Slice(&arr[0], 10) // equivalent to sss := (*[10]int)(unsafe.Pointer(&arr[0]))[:]
	ss[2] = 11
	ss[3] = 22
	ss = append(ss, 3, 4, 5)
	t.Logf("arr.type: %v", (reflect.TypeOf(arr))) // [2]int
	t.Logf("ss.type: %v", reflect.TypeOf(ss))
	t.Logf("ss: %v", ss) // ss: [9 8 11 22 0 0 0 0 0 0 3 4 5]
	t.Logf("ss.pointer: %p", &ss)
	t.Logf("ss[0].pointer: %p", &ss[0])
	t.Logf("len(ss): %v", len(ss))
	t.Logf("cap(ss): %v", cap(ss))
	ssrefhead := (*reflect.SliceHeader)(unsafe.Pointer(&ss))
	t.Logf("ssrefhead.pointer: %p", ssrefhead)                   // same as ss.pointer
	t.Logf("ssrefhead.Data: %p", unsafe.Pointer(ssrefhead.Data)) // same as ss[0].pointer
	t.Logf("ssrefhead.Len: %v", ssrefhead.Len)
	t.Logf("ssrefhead.Cap: %v", ssrefhead.Cap)
}

func TestUnsafeOffsetof(t *testing.T) {
	var strt STRUCT

	t.Log(unsafe.Pointer(&strt))
	// 0xc0000182e0, address(uintptr(HEX)) of strt, WARINING: &strt == &strt.i

	t.Log(uintptr(unsafe.Pointer(&strt)))
	// 824633819872, address(uintptr(DEC)) of strt

	/* unsafe.Pointer(&strt.i) == unsafe.Pointer(&strt) */
	t.Log(unsafe.Pointer(&strt.i))
	// 0xc0000182e0, address(unsafe.Pointer(HEX)) of strt.i

	t.Log(uintptr(unsafe.Pointer(&strt)) + unsafe.Offsetof(strt.i))
	// 824633819872, address(uintptr(DEC)) of strt.i

	t.Log(unsafe.Offsetof(strt.i)) // 0

	t.Log(unsafe.Pointer(&strt.j))
	// 0xc0000182e8, address(unsafe.Pointer(HEX)) of strt.j

	t.Log(uintptr(unsafe.Pointer(&strt)) + unsafe.Offsetof(strt.j))
	// 824633819880, address(uintptr(DEC)) of strt.j

	t.Log(unsafe.Offsetof(strt.j))
	// 8 == distance between start of strt and start of strt.j == size of strt.i

	u := uintptr(unsafe.Pointer(&strt))
	offset := unsafe.Offsetof(strt.j)
	t.Log(u + offset) // 0xc0000182e8
}

func TestUnsafeSizeof(t *testing.T) {
	var a int64
	var strt STRUCT
	var strt2 STRUCT2
	var strt3 STRUCT3
	arr := [10]uint16{}

	// func Sizeof(x ArbitraryType) uintptr
	// return size in bytes on memory of variable x
	// in case of struct variable, size comes with alignment
	// reference type (pointer) variable will be calculated in poninter variable size(system arch size)

	t.Log(unsafe.Sizeof(a))   // 8
	t.Log(unsafe.Sizeof(arr)) // 20

	t.Log(unsafe.Sizeof(strt))  // 24 with padding
	t.Log(unsafe.Sizeof(strt2)) // 24 with padding
	t.Log(unsafe.Sizeof(strt3)) // 8 with padding

	t.Log(unsafe.Sizeof(strt.i)) // 8
	t.Log(unsafe.Sizeof(strt.j)) // 4

	t.Log(unsafe.Sizeof(strt.k))  // 8
	t.Log(unsafe.Sizeof(strt2.k)) // 8
}

func TestUnsafeAlignof(t *testing.T) {
	var a int64
	var strt STRUCT
	var strt2 STRUCT2
	var strt3 STRUCT3

	// func Alignof(x ArbitraryType) uintptr
	// if x is struct variable) size of max_size_field which aligned_struct_size(with packing)%field_size==0
	// if x is non-struct variable) size of variable == unsafe.Sizeof(x)
	t.Log(reflect.TypeOf(a).Align()) // 8, non-struct variable
	t.Log(unsafe.Alignof(a))         // 8, non-struct variable
	t.Log(unsafe.Alignof(strt.i))    // 8, non-struct variable
	t.Log(unsafe.Alignof(strt.j))    // 4, non-struct variable
	t.Log(unsafe.Alignof(strt))      // 8, struct variable
	t.Log(unsafe.Alignof(strt2))     // 8, struct variable
	t.Log(unsafe.Alignof(strt3))     // 4, struct variable

	var un UNEFFICIENT
	t.Log(unsafe.Sizeof(un))  // 32
	t.Log(unsafe.Alignof(un)) // 8
}
