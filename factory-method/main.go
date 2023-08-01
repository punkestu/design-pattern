package main

import "github.com/punkestu/design-pattern/factory-method/lib"

func main() {
	var factory = lib.Factory{}
	product1 := factory.Produce("variant_1")
	product2 := factory.Produce("variant_2")
	println("first product description  :", product1.Describe())
	println("second product description :", product2.Describe())
}
