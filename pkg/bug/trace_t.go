package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	x := 100
	fmt.Printf("x addr:%v\n", &x)

	go func() {
		defer wg.Done()
		for {
			x += 1
		} }()

	go func() {
		defer wg.Done()
		for {
			x += 100
		} }()

	wg.Wait()
}

// GOMAXPROCS=2 go run -race main.go

