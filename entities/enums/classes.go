package enums

type Classes string

const (
	Knight    Classes = "KNIGHT"
	Ranger    Classes = "RANGER"
	Paladin   Classes = "PALADIN"
	Rogue     Classes = "ROGUE"
	Barbarian Classes = "BARBARIAN"
)

func (name Classes) String() string {
	return string(name)
}
