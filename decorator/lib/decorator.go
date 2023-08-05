package lib

import "fmt"

type Wrappee interface {
	Exec()
}

type Base struct{}

func (b Base) Exec() {
	fmt.Println("exec base wrappee")
}
