package methods

type Method1 struct{}

func (m Method1) Mod(data string) int {
	if data == "hello" {
		return 1
	}
	return 100
}

type Method2 struct{}

func (m Method2) Generate(data string) string {
	if data == "world" {
		return "yes"
	}
	return "no"
}
