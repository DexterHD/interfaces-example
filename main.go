package main

import (
	"fmt"

	"github.com/DexterHD/interfaces-example/animal"
	"github.com/DexterHD/interfaces-example/circus"
)

func main() {
	t := &circus.Tamer{}
	t2 := &circus.Tamer{}
	d := &animal.Dog{}
	c := &animal.Cat{}

	fmt.Println(t.Praise(d))  // woof
	fmt.Println(t.Praise(c))  // meow!
	fmt.Println(t.Praise(t2)) // Error: *circus.Tamer does not implement circus.Animal
}
