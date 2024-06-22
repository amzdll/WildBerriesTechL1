package main

import "fmt"

type Human struct {
	name string
}

func (h Human) Greet() {
	fmt.Printf("Hi, I'm %s!", h.name)
}

type Action struct {
	Human
}

func main() {
	act := Action{Human{"James"}}
	act.Greet()
}
