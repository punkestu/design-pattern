package lib

import (
	"github.com/punkestu/design-pattern/adapter/lib/object"
	"strconv"
)

type A interface {
	GetStringComponent() string
}

type BtoA struct {
	B *object.B
}

func (b BtoA) GetStringComponent() string {
	return strconv.Itoa(b.B.GetIntComponent())
}
