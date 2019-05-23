package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	var sliceNil []int
	sliceMakeNil := make([]int, 0, 0)

	// 是否为 nil
	fmt.Printf("sliceNil is nil: %+v\n", sliceNil == nil)
	fmt.Printf("sliceMakeNil is nil: %+v\n", sliceMakeNil == nil)
	fmt.Println()

	// size
	fmt.Printf("sliceNil size: %+v\n", unsafe.Sizeof(sliceNil))
	fmt.Printf("sliceMakeNil size: %+v\n", unsafe.Sizeof(sliceMakeNil))
	fmt.Println()

	// all value
	pSliceNil := uintptr(unsafe.Pointer(&sliceNil))
	fmt.Printf("sliceNil v[0]:%v, v[1]:%v, v[2]:%v\n",
		*(*uint64)(unsafe.Pointer(pSliceNil)),
		*(*uint64)(unsafe.Pointer(pSliceNil + 8)),
		*(*uint64)(unsafe.Pointer(pSliceNil + 16)))


	pSliceMakeNil := uintptr(unsafe.Pointer(&sliceMakeNil))
	fmt.Printf("sliceMakeNil v[0]:%v, v[1]:%v, v[2]:%v\n",
		*(*uint64)(unsafe.Pointer(pSliceMakeNil)),
		*(*uint64)(unsafe.Pointer(pSliceMakeNil + 8)),
		*(*uint64)(unsafe.Pointer(pSliceMakeNil + 16)))
	fmt.Println()

	// how is nil
	*(*uint64)(unsafe.Pointer(pSliceNil + 16)) = 10
	*(*uint64)(unsafe.Pointer(pSliceNil + 8)) = 10
	fmt.Printf("sliceNil v[0]:%v, v[1]:%v, v[2]:%v\n",
		*(*uint64)(unsafe.Pointer(pSliceNil)),
		*(*uint64)(unsafe.Pointer(pSliceNil + 8)),
		*(*uint64)(unsafe.Pointer(pSliceNil + 16)))
	fmt.Printf("sliceNil is nil: %+v\n", sliceNil == nil)
	fmt.Println()


	*(*uint64)(unsafe.Pointer(pSliceNil + 16)) = 0
	*(*uint64)(unsafe.Pointer(pSliceNil + 8)) = 0
	*(*uint64)(unsafe.Pointer(pSliceNil)) = 1
	fmt.Printf("sliceNil v[0]:%v, v[1]:%v, v[2]:%v\n",
		*(*uint64)(unsafe.Pointer(pSliceNil)),
		*(*uint64)(unsafe.Pointer(pSliceNil + 8)),
		*(*uint64)(unsafe.Pointer(pSliceNil + 16)))
	fmt.Printf("sliceNil is nil: %+v\n", sliceNil == nil)
	fmt.Println()

	// TODO
	// 1. slice how is nil?  first point != 0 ?
	// 源码地址: src/reflect/value:1032

	// 2. why make([]int, 0, 0) first point != 0?
	a := make([]int, 0, 0)
	b := make([]string, 0, 0)
	pA := uintptr(unsafe.Pointer(&a))
	pB := uintptr(unsafe.Pointer(&b))
	fmt.Printf("Pa:%v, Pb:%v\n", *(*uint64)(unsafe.Pointer(pA)),  *(*uint64)(unsafe.Pointer(pB)))
	fmt.Println()
	// 源码地址: src/runtime/slice.go:34

	// 3. slice == nil is good? len(slice) == 0 is ok?

}
