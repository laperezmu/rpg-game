package enums

type Races string

const (
	Teofrastos Races = "TEOFRASTOS"
	Argones    Races = "ARGONES"
	Eldenes    Races = "ELDENES"
	Alabrastos Races = "ALABRASTOS"
)

func (name Races) String() string {
	return string(name)
}
