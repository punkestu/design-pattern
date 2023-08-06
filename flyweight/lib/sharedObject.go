package lib

type SharedObject struct {
	ComponentS1 []string
	ComponentS2 []int
}

func NewSharedObject(component1 []string, component2 []int) *SharedObject {
	return &SharedObject{
		ComponentS1: component1,
		ComponentS2: component2,
	}
}
