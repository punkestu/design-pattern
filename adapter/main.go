package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/adapter/lib"
	"github.com/punkestu/design-pattern/adapter/lib/object"
)

func getStringComponentFromA(a lib.A) {
	fmt.Println(a.GetStringComponent())
}

func main() {
	a := object.A{Component: "hello"}
	println("From Object A")
	getStringComponentFromA(a)

	b := object.B{Component: 1000}
	println("From Object B")
	// getStringComponentFromA(b) will not work
	adapterB := lib.BtoA{B: &b}
	getStringComponentFromA(adapterB) // will work
}
