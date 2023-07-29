package lib

import "github.com/punkestu/design-pattern/factory-method/lib/product"

type Factory interface {
	Produce() product.Product
}

type FirstFactory struct{}

func (f FirstFactory) Produce() product.Product {
	return product.FirstProduct{}
}

type SecondFactory struct{}

func (s SecondFactory) Produce() product.Product {
	return product.SecondProduct{}
}
