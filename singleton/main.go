package main

import "github.com/punkestu/design-pattern/singleton/lib"

func main() {
	lib.InitSingleton("dobberman")
	lib.InitSingleton("minerva")
	lib.PrintName()
	lib.Run()
}
