package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Client %d\n", i)
		net.Dial("tcp", "127.0.0.1:2020")
	}
	time.Sleep(time.Duration(200) * time.Second)
}
