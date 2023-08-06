package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/composite/lib"
)

func main() {
	n1 := lib.NewNode(1)
	n11 := lib.NewNode(11)
	n111 := lib.NewNode(111)
	n112 := lib.NewNode(112)
	n12 := lib.NewNode(12)
	n121 := lib.NewNode(121)

	n1.RegChildren([]lib.INode{n11, n12})
	n11.RegChildren([]lib.INode{n111, n112})
	n12.RegChildren([]lib.INode{n121})

	fmt.Println(n1.GetBody())
}
