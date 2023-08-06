package main

import (
	"fmt"
	"github.com/punkestu/design-pattern/flyweight/lib"
	"runtime"
)

type HeavyWeight struct {
	ComponentS1 []string
	ComponentS2 []int
	ComponentM1 int
	ComponentM2 bool
}

func notFlyweight() {
	var mainObjs []HeavyWeight
	for i := 0; i < 1000; i++ {
		mainObjs = append(mainObjs, HeavyWeight{
			ComponentS1: []string{"hello", "world"},
			ComponentS2: []int{1, 2, 3},
			ComponentM1: 0,
			ComponentM2: false,
		})
	}
	fmt.Println("Not using flyweight:")
	PrintMemUsage()
}

func flyweight() {
	shared := lib.NewSharedObject(
		[]string{"hello", "world"},
		[]int{1, 2, 3},
	)
	var mainObjs []lib.MainObject
	for i := 0; i < 1000; i++ {
		mainObjs = append(mainObjs, lib.NewMainObject(shared))
	}
	fmt.Println("Using flyweight:")
	PrintMemUsage()
}

func main() {
	notFlyweight()
	flyweight()
}

var prev runtime.MemStats

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", float32(m.Alloc-prev.Alloc)/1000000)
	fmt.Printf("\tTotalAlloc = %v MiB", float32(m.TotalAlloc)/1000000)
	fmt.Println()
	prev = m
}
