package main

import (
	"unsafe"
	"time"
)

type data struct {
	x [1024 * 100]byte
}

func test() unsafe.Pointer {
	p := &data{}
	return unsafe.Pointer(p)
}
func main() {
	const N = 10000
	cache := new([N]unsafe.Pointer)
	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}
	}

// go build -o test && GODEBUG="gctrace=1" go run uintptr_t.go