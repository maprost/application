package texmodel

import "github.com/maprost/application/generator/lang"

type Skill struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Website struct {
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

type Education struct {
	Graduation string
	Institute  string
	Time       string
	Focus      string
}

type Index struct {
	Label           lang.Language
	Image           string // path
	Name            string
	Title           string
	Nationality     string
	Location        string
	Email           string
	Phone           string
	Websites        []Website
	ProfSkills      []Skill
	OtherProfSkills string
	Interest        string
	SoftSkills      string
	Language        []Language
	Hobbies         string
	AboutMe         string
	Experience      []Experience
	Education       []Education
	MainColor       string
	ShortVersion    bool
}