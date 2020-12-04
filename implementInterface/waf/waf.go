package waf

import (
	"fmt"
)

type Waf struct {
	name string
}

func (w *Waf) ExecuteLua(script string) {
	fmt.Println("Waf implement lua interface!")
}

func NewWaf() *Waf {
	return &Waf{name: "Waf"}
}
