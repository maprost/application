package texmodel

import (
	"github.com/maprost/application/generator/lang"
)

type Icon struct {
	Project   string
	Role      string
	TechStack string
	Document  string
}

type RatingLsa struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Language struct {
	Name  string
	Level string
}

type LeftSideAction struct {
	Type         int // 1: Skills + List, 2: List, 3: Language
	Label        string
	Ratings      []RatingLsa
	OtherRatings string
	List         string
	Languages    []Language
}

type Website struct {
	Icon string // path
	Url  string
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

type Index struct {
	// config
	MainColor    string
	SideColor    string
	HasProject   bool
	HasRole      bool
	HasTechStack bool
	Icon         Icon
	Label        lang.Language

	// main infos
	Image       string // path
	Name        string
	Title       string
	Nationality string
	Location    string
	Email       string
	Phone       string
	Websites    []Website

	// leftSide
	SideOneLeftSideAction []LeftSideAction
	SideTwoLeftSideAction []LeftSideAction

	// motivation
	AboutMeLabel string
	AboutMe      string

	// experience
	SideOneExperienceLabel string
	SideOneExperience      []Experience
	SideTwoExperienceLabel string
	SideTwoExperience      []Experience

	// education
	SideOneEducationLabel string
	SideOneEducation      []Education
	SideTwoEducationLabel string
	SideTwoEducation      []Education

	Attachment []string
}

func (x Index) HasSideTwo() bool {
	return len(x.SideTwoExperience) > 0 ||
		len(x.SideTwoEducation) > 0 ||
		len(x.SideTwoLeftSideAction) > 0
}
