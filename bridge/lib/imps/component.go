package imps

type ComponentA struct{}

func (c ComponentA) Get() string {
	return "Component A"
}

type ComponentB struct{}

func (c ComponentB) Get() string {
	return "Component B"
}
