package main

import "fmt"

func foo() *int {
	t := 3
	return &t;
}

func main() {
	x := foo()
	fmt.Println(*x)
}

//func main() {
//	a := new(int)
//	*a = 10
//	fmt.Println(*a)
//}


// go build -gcflags '-m -l' escape.go