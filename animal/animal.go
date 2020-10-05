package animal

// Dog implements dog animal.
type Dog struct{}

// Speaks dog do "woof".
func (d Dog) Speaks() string { return "woof" }

// IsTrained good boy.
func (d Dog) IsTrained() bool { return true }

// Jump yep, it can do it.
func (d Dog) Jump() string { return "jumps" }

// Sit also supported.
func (d Dog) Sit() string { return "sit" }

// Cat implements cat animal.
type Cat struct{}

// IsTrained nope folks. It's a cat.
func (c Cat) IsTrained() bool { return false }

// Speaks only meow...
func (c Cat) Speaks() string { return "meow!" }

// Jump no way...
func (c Cat) Jump() string { return "meow!!" }

// Sit I said no way!
func (c Cat) Sit() string { return "meow!!!" }
