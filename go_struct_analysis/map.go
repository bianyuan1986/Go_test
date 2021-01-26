package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")
	a := make(map[int]int)
	a[3]=4
	a[4]=5
	fmt.Println(a)
	b := make(map[int]int, 32)
	for i := 0; i < 32; i++ {
		b[i]=i
	}
	fmt.Println(b)
	c := make(map[int]string, 32)
	c[0]="2021"
	c[1]="2022"
	fmt.Println(c)
	d := *(&c)
	fmt.Println(d)
}
