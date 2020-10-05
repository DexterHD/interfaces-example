package main

import (
	"fmt"

	"github.com/idexter/interfaces-example/african"
	"github.com/idexter/interfaces-example/circus"
)

func main() {
	t := &circus.Tamer{}
	e := african.GetElephant()

	fmt.Println(t.Command(circus.ActVoice, e)) // "pawoo!"
	fmt.Println(t.Command(circus.ActJump, e))  // "o_O"
	fmt.Println(t.Command(circus.ActSit, e))   // "sit"
}
