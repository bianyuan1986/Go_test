package main

import (
	"fmt"
	"net"
	"net/http"
)

var stateNewCnt int
var stateActiveCnt int
var stateIdleCnt int
var stateHijackedCnt int
var stateClosedCnt int
var stateAliveCnt int

func connStateChangeHook(c net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		fmt.Println("State New!")
		stateNewCnt++
		stateAliveCnt++
		fmt.Printf("[Alive %d]\n", stateAliveCnt)

	case http.StateActive:
		fmt.Println("State Active!")
		stateActiveCnt++

	case http.StateIdle:
		fmt.Println("State Idle!")
		stateIdleCnt++

	case http.StateHijacked:
		fmt.Println("State Hijacked!")
		stateHijackedCnt++

	case http.StateClosed:
		fmt.Println("State Closed!")
		stateClosedCnt++
		stateAliveCnt--
		fmt.Printf("[Alive %d]\n", stateAliveCnt)
	}
}

func bianyuan_test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Congratulations!")
}

func NewServer(addr string) bool {
	mux := http.NewServeMux()
	mux.HandleFunc("/bianyuan_test", bianyuan_test)
	s := http.Server{Addr: addr, Handler: mux, ConnState: connStateChangeHook}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Listen failed!")
		return false
	}
	s.Serve(l)

	return true
}

func main() {
	NewServer("127.0.0.1:2020")
}
