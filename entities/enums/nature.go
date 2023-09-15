package enums

type Nature string

const (
	Rude     Nature = "RUDE"
	Cautious Nature = "CAUTIOUS"
	Active   Nature = "ACTIVE"
	Placid   Nature = "PLACID"
	Sassy    Nature = "SASSY"
	Neutral  Nature = "NEUTRAL"
)

func (name Nature) String() string {
	return string(name)
}
