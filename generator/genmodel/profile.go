package genmodel

type SkillID int

type Skill struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Experience struct {
	JobPosition string
	Company     string
	StartTime   string
	EndTime     string
	Description string
	TechStack   string
}

type Education struct {
	School      string
	StartTime   string
	EndTime     string
	Description string
}

type Profile struct {
	FirstName          string
	LastName           string
	Image              string // path
	Nationality        string
	Birthday           string
	Address            Address
	Email              string
	Website            []string          // url
	ProfessionalSkills map[SkillID]Skill // should contains all professional skills you have
	SoftSkills         map[SkillID]Skill // should contains all soft skills you have
	Interest           []string
	Hobbies            []string
	SignPath           string // path to the sign image
	Experience         []Experience
	Education          []Education
}
