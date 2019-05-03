package animal

// implementation of Animal
type Dog struct{}

func (d Dog) Speaks() string  { return "woof" }
func (d Dog) IsTrained() bool { return true }
func (d Dog) Jump() string    { return "jumps" }
func (d Dog) Sit() string     { return "sit" }

type Cat struct{}

func (c Cat) IsTrained() bool { return false }
func (c Cat) Speaks() string  { return "meow!" }
func (c Cat) Jump() string    { return "meow!!" }
func (c Cat) Sit() string     { return "meow!!!" }
