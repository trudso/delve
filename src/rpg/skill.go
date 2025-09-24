package rpg

type Skill struct {
	Name  string
	Level int

	Requirements Requirements
}

type Skills struct {
	Skills []Skill
}

func NewSkill(name string, level int) *Skill {
	return &Skill{
		Name:  name,
		Level: level,
	}
}


