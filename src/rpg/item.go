package rpg

type ItemType int
type ArmorType int
type WeaponType int

const (
	Armor ItemType = iota
	Weapon
	Amulet
	Ring
	Consumable
	Gold
)

const (
	Helmet ArmorType = iota
	Body
	Gloves
	Belt
	Legs
	Boots
	Shield
)

const (
	OneHandedMelee WeaponType = iota
	TwoHandedMelee
	Bow
	Quiver
)

type Item struct {
	Name string
	Type ItemType

	Requirements Requirements
	Modifiers    []Modifier
}
