package rpg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// character
type Character struct {
	Position   rl.Vector2
	Name       string
	Level      int
	Experience int

	Attributes Modifiable[Attributes]
	Skills     Modifiable[Skills]
	Equipment  Equipment
	Inventory  Inventory
}

func NewCharacter(name string, attributes Attributes) Character {
	return Character{
		Position:   rl.Vector2{},
		Name:       name,
		Attributes: Modifiable[Attributes]{Base: attributes, Current: attributes},
	}
}

func (c *Character) Refresh(modifiers []Modifier) {
	c.Attributes.Current = c.Attributes.Base
	for _, modifier := range modifiers {
		switch m := (modifier).(type) {
		case AttributesModifier:
			c.Attributes.Current = m.Modify(c.Attributes.Current)
		}
	}
}

type Attributes struct {
	Strength     int
	Agility      int
	Intelligence int
	Health       int
	Mana         int
}

type Requirements struct {
	Strength     int
	Agility      int
	Intelligence int
	Level        int
}

type Inventory struct {
	Items []Item
}

// items
type Equipment struct {
	Head      Item
	Body      Item
	Hands     Item
	Belt      Item
	Legs      Item
	Boots     Item
	RightHand Item
	LeftHand  Item
	Amulet    Item
	LeftRing  Item
	RightRing Item
}
