package lib

import "github.com/punkestu/design-pattern/factory-method/lib/product"

type Factory struct{}

func (f Factory) Produce(param int) product.Product {
	if param == 1 {
		return product.FirstProduct{}
	} else if param == 2 {
		return product.SecondProduct{}
	}
	return nil
}
