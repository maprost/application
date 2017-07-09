package genmodel

type SkillID int

type Skill struct {
	Name   string
	Rating int // Rating range [1,5]
}

type CV struct {
	Pic    string            // path
	Skills map[SkillID]Skill // should contains all skills you have
}
