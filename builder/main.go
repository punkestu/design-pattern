package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/builder/lib"
)

func main() {
	builder := lib.NewBuilder()
	building := builder.Generate()
	println("Before build")
	fmt.Println(building)
	builder.AttachComponent1("component 1")
	builder.AttachComponent2(100)
	builder.AttachComponent3(true)
	building = builder.Generate()
	println("After build")
	fmt.Println(building)
}
