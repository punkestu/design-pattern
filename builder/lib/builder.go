package lib

type Building struct {
	Component1 string
	Component2 int
	Component3 bool
}

type Builder struct {
	b *Building
}

func NewBuilder() *Builder {
	return &Builder{b: &Building{}}
}

func (b *Builder) Reset() {
	b.b = &Building{}
}

func (b *Builder) AttachComponent1(component string) {
	b.b.Component1 = component
}

func (b *Builder) AttachComponent2(component int) {
	b.b.Component2 = component
}

func (b *Builder) AttachComponent3(component bool) {
	b.b.Component3 = component
}

func (b *Builder) Generate() *Building {
	return b.b
}
