package main

import (
	"testing"
	"fmt"
	"time"
	"unsafe"
)

func TestLinkAdd(t *testing.T) {

	aH := unsafe.Pointer(&Node{})
	bH := unsafe.Pointer(&Node{})
	cH := unsafe.Pointer(&Node{})

	aL := &Link{head:aH, tail:aH}
	bL := &Link{head:bH, tail:bH}
	cL := &Link{head:cH, tail:cH}


	for i:=0; i<20;i++{
		go func(idx int) {
			fmt.Printf("idx:%d\n", idx)
			aL.LockAdd(idx)
			bL.noLockAdd(idx)
			cL.casLockAdd(idx)

		}(i)
	}

	time.Sleep(10 * time.Second)
	fmt.Printf("aL:%s\n", aL.String())
	fmt.Printf("bL:%s\n", bL.String())
	fmt.Printf("cL:%s\n", cL.String())


}