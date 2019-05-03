package circus

const (
	ActVoice = iota
)

type Speaker interface {
	Speaks() string
}

type Tamer struct{}

func (t *Tamer) Speaks() string { return "WAT?" }

func (t *Tamer) Command(action int, a Speaker) string {
	switch action {
	case ActVoice:
		return a.Speaks()
	}
	return ""
}

func Perform(a Speaker) string { return a.Speaks() }
