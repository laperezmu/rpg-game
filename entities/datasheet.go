package entities

import (
	"rpg-game/entities/contracts"
	"rpg-game/entities/enums"
)

type Datasheet struct {
	Name       string
	Age        int
	Background string
	Level      int
	Race       enums.Races
	Class      enums.Classes
	Element    enums.Element
	Nature     enums.Nature
	Attributes Attributes
	Inventory  []contracts.Object
	NPC        bool
}

type Attributes struct {
	Health    float64
	Energy    float64
	Attack    float64
	Defense   float64
	Evasion   float64
	Precision float64
	Dexterity float64
}
