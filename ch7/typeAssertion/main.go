package main

import "fmt"

type Dog struct {
	name string
}

func (Dog) run() {
	fmt.Printf("Dog is running\n")
}

type Cat struct {
	name string
}

func (Cat) run() {
	fmt.Printf("Cat is running\n")
}

func (Cat) miau() {
	fmt.Printf("Miau\n")
}

type Animal interface {
	run()
}

func main() {
	animals := []Animal{
		Dog{name: "boby"},
		Cat{name: "bigotes"},
	}

	for _, animal := range animals {
		animal.run()
		cat, ok := animal.(Cat)
		if ok {
			cat.miau()
		}
	}
}
