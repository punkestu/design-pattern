package main

import "github.com/punkestu/design-pattern/prototype/lib"

func main() {
	obj1 := lib.NewObject("obj1", 100)
	obj2 := obj1.Clone()
	println("Before modification")
	println("          |    Address   |   Value")
	print("From obj1 | ")
	obj1.Display()
	print("From obj2 | ")
	obj2.Display()

	// coba ganti ID nya
	obj2.(*lib.Object).ID = "obj2"
	println("After modification")
	println("          |    Address   |   Value")
	print("From obj1 | ")
	obj1.Display()
	print("From obj2 | ")
	obj2.Display()
}
