package main

import (
	"fmt"

	"github.com/DexterHD/interfaces-example/animal"
	"github.com/DexterHD/interfaces-example/circus"
)

func main() {
	d := &animal.Dog{}
	t := &circus.Tamer{}
	t2 := &circus.Tamer{}

	fmt.Println(t.Command(circus.ActVoice, d))  // woof
	fmt.Println(t.Command(circus.ActVoice, t2)) // WAT?
}
