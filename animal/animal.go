package animal

// implementation of Animal
type Dog struct{}

func (a Dog) Speaks() string  { return "woof" }
func (a Dog) IsTrained() bool { return true }
