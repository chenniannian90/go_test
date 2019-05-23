package main

import (
		_ "net/http/pprof"
		"net/http"
		"time"
)

func do()  {

	a := make([]int, 0, 10000)
	for i:=0; i<10000; i++{
		a = append(a, i)
	}

	sum := 0
	for _, v:=range a{
		sum += v
	}

}

func main() {



	go http.ListenAndServe("localhost:6060", nil)
	for {

		go do()
		time.Sleep(time.Second)

	}
}

// http://localhost:6060/debug/pprof/
// go tool pprof pprof heap

