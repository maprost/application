package genmodel

type SkillID int

type Skill struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Experience struct {
	JobPosition      string
	Company          string
	StartTime        string
	EndTime          string
	Description      string
	TechStack        string
	FutureExperience bool
}

func (e Experience) Empty() bool {
	return e.JobPosition == ""
}

type Education struct {
	Graduation string
	Institute  string
	StartTime  string
	EndTime    string
	Focus      string
	FinalGrade string
}

type Language struct {
	Name  string
	Level string
}

type Profile struct {
	FirstName             string
	LastName              string
	Image                 string // path
	Nationality           string
	Birthday              string
	Address               Address
	Email                 string
	Phone                 string
	Websites              []string          // url
	ProfessionalSkills    map[SkillID]Skill // should contains all professional skills you have
	SoftSkills            map[SkillID]Skill // should contains all soft skills you have
	Interest              []string
	Hobbies               []string
	SignPath              string // path to the sign image
	Experience            []Experience
	Education             []Education
	Language              []Language
	GeneralMotivationText string // your general motivation text, this text can contains tex elements
}
