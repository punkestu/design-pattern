package decorators

import (
	"fmt"
	"github.com/punkestu/design-pattern/decorator/lib"
)

type Wrapper1 struct {
	W lib.Wrappee
}

func (w Wrapper1) Exec() {
	fmt.Println("Exec wrapper1")
	w.DoSomething()
	w.W.Exec()
}

func (w Wrapper1) DoSomething() {
	fmt.Println("Wrapper1 do something")
}

type Wrapper2 struct {
	W lib.Wrappee
}

func (w Wrapper2) Exec() {
	fmt.Println("Exec wrapper2")
	w.DoSomething()
	w.W.Exec()
}

func (w Wrapper2) DoSomething() {
	fmt.Println("Wrapper2 do something")
}
