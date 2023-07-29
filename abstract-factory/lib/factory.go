package lib

type Factory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ProductA interface {
	Action1() string
	Action2() int
}
type ProductB interface {
	Action() bool
}
