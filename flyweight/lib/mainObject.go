package lib

import "fmt"

type MainObject struct {
	Shared      *SharedObject
	ComponentM1 int
	ComponentM2 bool
}

func NewMainObject(shared *SharedObject) MainObject {
	return MainObject{
		Shared:      shared,
		ComponentM1: 0,
		ComponentM2: false,
	}
}

func (m *MainObject) DoSomething() {
	fmt.Println(m)
}
