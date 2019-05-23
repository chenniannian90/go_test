package main

import (
	"fmt"
	"sync"
	"strings"
	"sync/atomic"
	"unsafe"
	"time"
)

type Node struct {
	value int
	next unsafe.Pointer
}

func (node Node)String() string {
	return fmt.Sprintf("Value:%d", node.value)
}


type Link struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	sync.Mutex
}

func (link Link)String() string{

	nodeStrList := make([]string, 0, 100)
	cur := (*Node)((*Node)(link.head).next)
	for cur != nil {
		nodeStrList = append(nodeStrList, cur.String())
		cur = (*Node)((*Node)(cur).next)
	}
	return strings.Join(nodeStrList, " ")
}


func (link *Link)noLockAdd(value int){

	node := Node{value:value, next:nil}
	tail := (*Node)(link.tail)
	tail.next = unsafe.Pointer(&node)
	time.Sleep(1 * time.Microsecond)
	link.tail = unsafe.Pointer(&node)


}

func (link *Link)LockAdd(value int){

	link.Mutex.Lock()
	defer link.Mutex.Unlock()

	node := Node{value:value, next:nil}
	tail := (*Node)(link.tail)
	tail.next = unsafe.Pointer(&node)
	time.Sleep(1 * time.Microsecond)
	link.tail = unsafe.Pointer(&node)
}

func (link *Link)casLockAdd(value int){

	node := Node{value:value, next:nil}
	// https://coolshell.cn/articles/8239.html 参考

	var tail *Node
	for{
		tail = (*Node)(link.tail)
		time.Sleep(1 * time.Microsecond)
		if ok:= atomic.CompareAndSwapPointer(&tail.next, nil, unsafe.Pointer(&node)); ok{
			break
		}
	}

	time.Sleep(1 * time.Microsecond)
	atomic.CompareAndSwapPointer(&link.tail, unsafe.Pointer(tail), unsafe.Pointer(&node))

}

func main()  {

	aH := unsafe.Pointer(&Node{})
	bH := unsafe.Pointer(&Node{})
	cH := unsafe.Pointer(&Node{})

	aL := &Link{head:aH, tail:aH}
	bL := &Link{head:bH, tail:bH}
	cL := &Link{head:cH, tail:cH}

	num := 20
	var waitGroup sync.WaitGroup
	waitGroup.Add(20*3)

	for i:=0; i<num;i++{
		go func(idx int) {
			defer waitGroup.Done()
			aL.LockAdd(idx)

		}(i)
		go func(idx int) {
			defer waitGroup.Done()
			bL.noLockAdd(idx)

		}(i)
		go func(idx int) {
			defer waitGroup.Done()
			cL.casLockAdd(idx)

		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("aL:%s\n", aL.String())
	fmt.Printf("bL:%s\n", bL.String())
	fmt.Printf("cL:%s\n", cL.String())
}