package enums

type Element string

const (
	Aqua   Element = "AQUA"
	Ignis  Element = "IGNIS"
	Ventum Element = "VENTUM"
	Terra  Element = "TERRA"
	Fulgur Element = "FULGUR"
	Umbra  Element = "UMBRA"
	Lux    Element = "LUX"
	Enton  Element = "ENTON"
)

func (name Element) String() string {
	return string(name)
}
