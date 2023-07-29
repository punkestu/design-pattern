package factory2

import (
	"github.com/punkestu/design-pattern/abstract-factory/lib"
	"github.com/punkestu/design-pattern/abstract-factory/lib/factory2/product"
)

type Factory2 struct{}

func (f Factory2) CreateProductA() lib.ProductA {
	return product.A{}
}

func (f Factory2) CreateProductB() lib.ProductB {
	return product.B{}
}
