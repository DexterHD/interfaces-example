package main

import (
	"fmt"

	"github.com/DexterHD/interfaces-example/animal"
	"github.com/DexterHD/interfaces-example/circus"
)

func main() {
	t := &circus.Tamer{}
	d := &animal.Dog{}

	fmt.Println(t.Command(circus.ActVoice, d)) // "woof"
	fmt.Println(t.Command(circus.ActJump, d))  // "jumps"
	fmt.Println(t.Command(circus.ActSit, d))   // "sit"

	t2 := &circus.Tamer{}
	c := &animal.Cat{}
	fmt.Println(t2.Command(circus.ActVoice, c)) // "panic: Sorry but this animal doesn't understand your commands"
	fmt.Println(t2.Command(circus.ActJump, c))  // ...
	fmt.Println(t2.Command(circus.ActSit, c))   // ...
}
