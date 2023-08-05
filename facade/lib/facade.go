package lib

import "github.com/punkestu/design-pattern/facade/lib/methods"

type Methods string

const (
	Method1 Methods = "1"
	Method2 Methods = "2"
)

type Payload struct {
	Method Methods
	Value  string
}

type Facade struct {
	method1 methods.Method1
	method2 methods.Method2
}

func NewFacade() Facade {
	return Facade{
		method1: methods.Method1{},
		method2: methods.Method2{},
	}
}

func (f Facade) Transform(payload Payload) string {
	if payload.Method == Method1 {
		if f.method1.Mod(payload.Value) != 1 {
			return "Not Hello"
		}
		return "Is Hello"
	} else if payload.Method == Method2 {
		if f.method2.Generate(payload.Value) == "no" {
			return "Is World"
		}
		return "Not World"
	}
	return ""
}
