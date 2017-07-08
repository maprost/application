package genmodel

type SkillID int

type Skill struct {
	ID     SkillID
	Name   string
	Rating int // Rating range [1,5]
}

type CV struct {
	Pic string // path
	Skills []Skill // should contains all skills you have
}
