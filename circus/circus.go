package circus

const (
	ActVoice = iota
)

type Speaker interface {
	Speaks() string
}

type Animal interface {
	Speaker
}

type Tamer struct{}

func (t *Tamer) Speaks() string { return "WAT?" }

func (t *Tamer) Command(action int, a Animal) string {
	switch action {
	case ActVoice:
		return a.Speaks()
	}
	return ""
}

func Perform(a Speaker) string { return a.Speaks() }
