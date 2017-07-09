package texmodel

type Skill struct {
	Name   string
	Rating int // Rating range [1,5]
}

type Websites struct {
	Icon string // path
	Url  string
}

type CV struct {
	Pic      string // path
	Name     string
	Title    string
	Address  string
	Email    string
	Phone    string
	Websites []Websites
	Skills   []Skill
}
