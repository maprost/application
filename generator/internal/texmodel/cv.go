package texmodel

type Skill struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Websites struct {
	Icon string // path
	Url  string
}

type Language struct {
	Name  string
	Level string
}

type Experience struct {
	Company     string
	Position    string
	Time        string
	Description string
	Tech        string
}

type CV struct {
	Pic         string // path
	Name        string
	Title       string
	Nationality string
	Location    string
	Email       string
	Phone       string

	Websites    []Websites
	Skills      []Skill
	OtherSkills string
	Language    []Language
	Experience  []Experience
}
