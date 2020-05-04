// ex2/reg.go
package ex1

import (
	"fmt"

	"../base"
)

type Class2 struct {
}

func (c *Class2) Do() {
	fmt.Println("class2")
}

func init() {
	base.Register("Class2", func() base.Class {
		return new(Class2)
	})
}
