package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/facade/lib"
)

func main() {
	facade := lib.Facade{}
	res := facade.Transform(
		lib.Payload{
			Method: lib.Method1,
			Value:  "world",
		},
	)
	fmt.Println(res)
}
