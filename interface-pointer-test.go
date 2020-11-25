package main

import (
	"fmt"
)

type OutInf interface {
	Output() bool
}

type OutSt struct {
	name string
}

/*实现了OutInf接口方法*/
func (s *OutSt) Output() bool {
	fmt.Printf("%s\n", s.name)

	return true
}

func MyPrint(out interface{}) {
	a := out.(OutInf)
	a.Output()

	c := out.(*OutSt)
	c.Output()

	b := out.(OutSt)
	b.Output()
}

func main() {
	outP := &OutSt{name: "Tiger"}
	MyPrint(outP)

	var outS OutSt
	outS.name = "baohu"
	MyPrint(&outS)

	fmt.Println("End!")
}
