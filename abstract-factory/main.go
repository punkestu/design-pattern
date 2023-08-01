package main

import (
	"github.com/punkestu/design-pattern/abstract-factory/lib/factory1"
	"github.com/punkestu/design-pattern/abstract-factory/lib/factory2"
)

func main() {
	factory_1 := factory1.Factory1{}
	factory_2 := factory2.Factory2{}
	productA1 := factory_1.CreateProductA("variant_1")
	productB1 := factory_1.CreateProductB("variant_2")
	productA2 := factory_2.CreateProductA("variant_2")
	productB2 := factory_2.CreateProductB("variant_1")
	println("from factory_1")
	println(productA1.Action1(), productA1.Action2())
	println(productB1.Action())
	println()
	println("from factory_2")
	println(productA2.Action1(), productA2.Action2())
	println(productB2.Action())
}
