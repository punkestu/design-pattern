package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/bridge/lib/imps"
)

func main() {
	comp1 := imps.ComponentA{}
	comp2 := imps.ComponentB{}
	obj1 := imps.ObjectA{Comp: comp1}
	obj2 := imps.ObjectA{Comp: comp2}
	obj3 := imps.ObjectB{Comp: comp1}
	obj4 := imps.ObjectB{Comp: comp2}
	fmt.Println(obj1.Describe())
	fmt.Println(obj2.Describe())
	fmt.Println(obj3.Describe())
	fmt.Println(obj4.Describe())
}
