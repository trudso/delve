package rpg

type Modifiable[T any] struct {
	Base    T
	Current T
}

type Modifier any

type AttributesModifier interface {
	Modify(attr Attributes) Attributes
}

type AttributesModification struct {
	Modification Attributes
}

func (m AttributesModification) Modify(attr Attributes) Attributes {
	return Attributes{
		Strength:     m.Modification.Strength + attr.Strength,
		Agility:      m.Modification.Agility + attr.Agility,
		Intelligence: m.Modification.Intelligence + attr.Intelligence,
	}
}
