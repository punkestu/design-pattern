package product

type A struct{}

func (a A) Action1() string {
	return "this is product A from factory 1"
}

func (a A) Action2() int {
	return 100
}
