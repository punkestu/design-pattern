package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/decorator/lib"
	"github.com/punkestu/design-pattern/decorator/lib/decorators"
)

func main() {
	base := lib.Base{}
	base.Exec()
	fmt.Println("==========")
	deco1 := decorators.Wrapper1{W: base}
	deco1.Exec()
	fmt.Println("==========")
	deco2 := decorators.Wrapper2{W: deco1}
	deco2.Exec()
}
