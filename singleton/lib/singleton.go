package lib

import "log"

type singleton struct {
	Name string
}

var instance *singleton

func InitSingleton(name string) {
	if instance == nil {
		instance = &singleton{name}
		log.Println("object initialized")
		return
	}
	log.Println("object cannot reinitialize")
}

func PrintName() {
	println(instance.Name)
}

func Run() {
	println(instance.Name, "is Running")
}
