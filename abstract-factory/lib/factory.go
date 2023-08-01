package lib

type Factory interface {
	CreateProductA(variant string) ProductA
	CreateProductB(variant string) ProductB
}

type ProductA interface {
	Action1() string
	Action2() int
}
type ProductB interface {
	Action() bool
}
