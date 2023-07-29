package main

import "github.com/punkestu/design-pattern/factory-method/lib"

func main() {
	var factory1 lib.Factory = lib.FirstFactory{}
	var factory2 lib.Factory = lib.SecondFactory{}
	product1 := factory1.Produce()
	product2 := factory2.Produce()
	println("first product description  :", product1.Describe())
	println("second product description :", product2.Describe())
}
