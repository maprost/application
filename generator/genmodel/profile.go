package genmodel

type SkillID int

type Skill struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Profile struct {
	FirstName       string
	LastName        string
	Nationality     string
	Birthday        string
	Address         Address
	Website         []string          // url
	TechnicalSkills map[SkillID]Skill // should contains all technical skills you have
	SoftSkills      map[SkillID]Skill // should contains all soft skills you have
	Interest        []string
	Hobbies         []string
	SignPath        string // path to the sign image
}
