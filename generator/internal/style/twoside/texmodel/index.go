package texmodel

import (
	"github.com/maprost/application/generator/genmodel"
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

type RangeLsa struct {
	Name            string
	Min             float64 // range [0,1]
	MinLabel        string
	MinString       string  // needed for .tex
	Max             float64 // range [0,1]
	MaxLabel        string
	MaxString       string // needed for .tex
	DeltaMaxMin     string
	DeltaMaxMinHalf string
	DeltaMaxFull    string
	SingleLabel     string
}

type LeftSideAction struct {
	Type         int // 1: Skills + List, 2: List, 3: Language
	Label        string
	Ratings      []RatingLsa
	OtherRatings string
	List         string
	Languages    []Language
	Range        []RangeLsa
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

type Publication struct {
	Title         string
	Publisher     string
	Time          string
	Content       string
	DocumentLinks []string
}

type Award struct {
	Title         string
	Institute     string
	Time          string
	Content       string
	DocumentLinks []string
}

type Index struct {
	// config
	MainColor        string
	SideColor        string
	HasDocumentLinks bool
	HasProject       bool
	HasRole          bool
	HasTechStack     bool
	Icon             Icon
	Label            lang.Language

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

	// about me
	AboutMeLabel string
	AboutMe      string

	// motivation
	MyMotivationLabel string
	MyMotivation      string

	// main question
	MainQuestionLabel string
	MainQuestion      string

	//TODO money amount
	MoneyAmountLabel string
	MoneyAmounts     []genmodel.LeftSideAction

	//TODO time amount
	TimeAmountLabel string
	TimeAmounts     []genmodel.LeftSideAction

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

	// publication
	SideOnePublicationLabel string
	SideOnePublication      []Publication
	SideTwoPublicationLabel string
	SideTwoPublication      []Publication

	// award
	SideOneAwardLabel string
	SideOneAward      []Award
	SideTwoAwardLabel string
	SideTwoAward      []Award

	Attachment []string
}

func (x Index) HasSideTwo() bool {
	return len(x.SideTwoExperience) > 0 ||
		len(x.SideTwoEducation) > 0 ||
		len(x.SideTwoLeftSideAction) > 0
}
