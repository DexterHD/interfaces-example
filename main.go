package main

import (
	"fmt"

	"github.com/DexterHD/interfaces-example/animal"
	"github.com/DexterHD/interfaces-example/circus"
)

func main() {
	d := &animal.Dog{}
	fmt.Println(circus.Perform(d)) // woof
}
