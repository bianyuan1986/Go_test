package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

/*返回值可以被命名，所以该函数返回后的返回值不为空，即err*/
func testError() (err error) {
	err = nil

	defer func() {
		err = errors.New("Error occured!")
	}()

	return
}

func main() {
	err := testError()
	if err == nil {
		fmt.Println("Error is nil!")
	} else {
		fmt.Println("Error is not nil!")
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("Client %d\n", i)
		net.Dial("tcp", "127.0.0.1:2020")
	}
	time.Sleep(time.Duration(1) * time.Second)
}
