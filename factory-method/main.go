package main

import "github.com/punkestu/design-pattern/factory-method/lib"

func main() {
	factory := lib.Factory{}
	product1 := factory.Produce(1)
	product2 := factory.Produce(2)
	println("first product description  :", product1.Describe())
	println("second product description :", product2.Describe())
}
