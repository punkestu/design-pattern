package product

type Product interface {
	Describe() string
}

type FirstProduct struct{}

func (p FirstProduct) Describe() string {
	return "this is first product"
}

type SecondProduct struct{}

func (s SecondProduct) Describe() string {
	return "this is second product"
}
