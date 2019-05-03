package circus

type Speaker interface {
	Speaks() string
}

type Tamer struct{}

func (t *Tamer) Speaks() string { return "WAT?" }

func Perform(a Speaker) string { return a.Speaks() }
