package main

import (
	"fmt"
)

type Action struct {
	Human
}

type Human struct {
	firstName string
	lastName  string
}

func (h Human) Say() {
	fmt.Printf("%v %v говорит", h.firstName, h.lastName)
}

func main() {
	action := &Action{Human: Human{"Владислав", "Синявский"}}
	action.Say()
}