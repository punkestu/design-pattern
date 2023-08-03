package object

type A struct {
	Component string
}

func (a A) GetStringComponent() string {
	return a.Component
}
