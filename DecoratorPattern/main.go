package main

import "fmt"

func main() {
	expresso := Expresso{}
	hola := NewMocha(&expresso)
	fmt.Printf("The Cost is: %d\n", hola.getCost())
	fmt.Printf("The Description is: %s\n", hola.getDescription())
}
