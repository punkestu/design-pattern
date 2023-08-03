package object

type B struct {
	Component int
}

func (b B) GetIntComponent() int {
	return b.Component
}
