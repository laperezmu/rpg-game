package entities

type Weapon struct {
	Name   string
	Damage float64
}

func (w Weapon) GetName() string {
	return w.Name
}
