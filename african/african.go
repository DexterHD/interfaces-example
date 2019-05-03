package african

import "github.com/DexterHD/interfaces-example/circus"

type Elephant struct{}

func (e Elephant) Speaks() string  { return "pawoo!" }
func (e Elephant) Jump() string    { return "o_O" }
func (e Elephant) Sit() string     { return "sit" }
func (e Elephant) IsTrained() bool { return true }

func GetElephant() circus.Animal { return &Elephant{} }
