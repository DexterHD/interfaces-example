package circus

const (
	ActVoice = iota
	ActSit
	ActJump
)

type Speaker interface {
	Speaks() string
}

type Animal interface {
	Speaker
	IsTrained() bool
	Jump() string
	Sit() string
}

type Tamer struct{}

func (t *Tamer) Speaks() string { return "WAT?" }

func (t *Tamer) Command(action int, a Animal) string {
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

func Perform(a Speaker) string { return a.Speaks() }
