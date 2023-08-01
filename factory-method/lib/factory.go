package lib

import "github.com/punkestu/design-pattern/factory-method/lib/product"

type Factory struct{}

func (f Factory) Produce(variant string) product.Product {
	if variant == "variant_1" {
		return product.FirstProduct{}
	} else if variant == "variant_2" {
		return product.SecondProduct{}
	} else {
		return nil
	}
}
