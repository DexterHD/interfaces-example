package circus

const (
	// ActVoice make some noise.
	ActVoice = iota
	// ActSit it won't stand up.
	ActSit
	// ActJump if you want you pet to fly.
	ActJump
)

// Speaker describes speaker ADT.
type Speaker interface {
	Speaks() string
}

// Animal described Animal ADT.
type Animal interface {
	Speaker
	IsTrained() bool
	Jump() string
	Sit() string
}

// Tamer implements circus tamer.
type Tamer struct{}

// Speaks yes, he/she can speak.
func (t *Tamer) Speaks() string { return "WAT?" }

// Command to do some action to an abstract Animal.
func (t *Tamer) Command(action int, a Animal) string {
	if !a.IsTrained() {
		panic("Sorry but this animal doesn't understand your commands")
		return ""
	}

	switch action {
	case ActVoice:
		return a.Speaks()
	case ActSit:
		return a.Sit()
	case ActJump:
		return a.Jump()
	}
	return ""
}

// Praise we need praise our pets.
func (t *Tamer) Praise(a Animal) string {
	return a.Speaks()
}

// Perform performs Speak action.
func Perform(a Speaker) string { return a.Speaks() }
