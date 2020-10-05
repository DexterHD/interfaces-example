package african

import "github.com/idexter/interfaces-example/circus"

// Elephant described elephant concrete type
type Elephant struct{}

// Speaks elephant can speak.
func (e Elephant) Speaks() string { return "pawoo!" }

// Jump elephant pretty puzzled about jumping.
func (e Elephant) Jump() string { return "o_O" }

// Sit yep it can sit.
func (e Elephant) Sit() string { return "sit" }

// IsTrained we have trained elephant.
func (e Elephant) IsTrained() bool { return true }

// GetElephant creates Animal which is Elephant in fact.
func GetElephant() circus.Animal { return &Elephant{} }
