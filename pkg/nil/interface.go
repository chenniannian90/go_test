package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

type ITest interface {
	test()
}

type Test struct {

}

func (t Test) test()  {

}

func main()  {


	var test1, test2, test3 ITest
	test2 = ITest(nil)
	test3 = (*Test)(nil)

	// 是否为 nil
	fmt.Printf("test1 is nil: %+v\n", test1 == nil)
	fmt.Printf("test2 is nil: %+v\n", test2 == nil)
	fmt.Printf("test3 is nil: %+v\n", test3 == nil)
	fmt.Println()

	// size
	fmt.Printf("test1 size: %+v\n", unsafe.Sizeof(test1))
	fmt.Printf("test2 size: %+v\n", unsafe.Sizeof(test2))
	fmt.Printf("test3 size: %+v\n", unsafe.Sizeof(test3))
	fmt.Println()

	// all value
	p1 := uintptr(unsafe.Pointer(&test1))
	fmt.Printf("test1 v[0]:%v, v[1]:%v\n",
		*(*uint64)(unsafe.Pointer(p1)),
		*(*uint64)(unsafe.Pointer(p1 + 8)))

	p2 := uintptr(unsafe.Pointer(&test2))
	fmt.Printf("test2 v[0]:%v, v[1]:%v\n",
		*(*uint64)(unsafe.Pointer(p2)),
		*(*uint64)(unsafe.Pointer(p2 + 8)))

	p3 := uintptr(unsafe.Pointer(&test3))
	fmt.Printf("test3 v[0]:%v, v[1]:%v\n",
		*(*uint64)(unsafe.Pointer(p3)),
		*(*uint64)(unsafe.Pointer(p3 + 8)))
	fmt.Println()

	// how is nil
	*(*uint64)(unsafe.Pointer(p1 + 8)) = 10
	fmt.Printf("test1 v[0]:%v, v[1]:%v\n",
		*(*uint64)(unsafe.Pointer(p1)),
		*(*uint64)(unsafe.Pointer(p1 + 8)))
	fmt.Printf("test1 is nil: %+v\n", test1 == nil)

	*(*uint64)(unsafe.Pointer(p1 + 8)) = 0
	*(*uint64)(unsafe.Pointer(p1)) = 10
	fmt.Printf("test1 v[0]:%v, v[1]:%v\n",
		*(*uint64)(unsafe.Pointer(p1)),
		*(*uint64)(unsafe.Pointer(p1 + 8)))
	fmt.Printf("test1 is nil: %+v\n", test1 == nil)
	fmt.Println()

	reflect.ValueOf(test3).IsNil()

	// TODO
	// 1. interface how is nil?  first point != 0 ?
	// 源码地址: src/reflect/value:1032


	// 2. interface == nil is good? reflect.ValueOf(test3).IsNil() is ok?

}
