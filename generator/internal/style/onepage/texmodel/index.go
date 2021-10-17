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
	Company       string
	Position      string
	Time          string
	Description   string
	Project       string // please add here your pet projects
	Role          string // please add here more roles you did
	Tech          string
	DocumentLinks []string
}

type Education struct {
	Graduation    string
	FinalGrade    string
	Institute     string
	Time          string
	Focus         string
	DocumentLinks []string
}

type Icon struct {
	Project   string
	Role      string
	TechStack string
	Document  string
}

type Index struct {
	Icon            Icon
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
	SideColor       string
	ShortVersion    bool
	Attachment      []string
	HasFocus        bool
}
