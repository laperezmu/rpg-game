package entities

type Armor struct {
	Name       string
	Resistance float64
}

func (a Armor) GetName() string {
	return a.Name
}
