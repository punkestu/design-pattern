package factory1

import (
	"github.com/punkestu/design-pattern/abstract-factory/lib"
	"github.com/punkestu/design-pattern/abstract-factory/lib/factory1/product"
)

type Factory1 struct{}

func (f Factory1) CreateProductA(variant string) lib.ProductA {
	return product.A{}
}

func (f Factory1) CreateProductB(variant string) lib.ProductB {
	return product.B{}
}
