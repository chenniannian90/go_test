package main

import (
	"testing"
	"fmt"
)


// 1.  mock 函数
var sayHello =  func () string{
	return "Hello"
}


func TestMockFunc(t *testing.T) {

	oldSayHello := sayHello
	defer func() {
		sayHello = oldSayHello
	}()

	sayHello = func() string {
		return "Mocked Hello"
	}
	fmt.Printf("sayHello return:%+v\n", sayHello())
}

