package lib

import "fmt"

type Prototype interface {
	Display()
	Clone() Prototype
}

type Object struct {
	ID        string
	component int
}

func NewObject(ID string, component int) *Object {
	return &Object{
		ID:        ID,
		component: component,
	}
}

func (o *Object) Display() {
	fmt.Println(&o, "|", o.ID, "|", o.component)
}

func (o *Object) Clone() Prototype {
	return &Object{
		ID:        o.ID,
		component: o.component,
	}
}
