package imps

type Component interface {
	Get() string
}

type ObjectA struct {
	Comp Component
}

func (o ObjectA) Describe() string {
	return "Object A with component: " + o.Comp.Get()
}

type ObjectB struct {
	Comp Component
}

func (o ObjectB) Describe() string {
	return "Object B with component: " + o.Comp.Get()
}
