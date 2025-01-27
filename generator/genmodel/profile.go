package genmodel

import "github.com/maprost/application/generator/lang"

type ID int

type LeftSideAction struct {
	Id          ID
	Name        lang.TranslationMap
	SingleLabel lang.TranslationMap
	Rating      int     // Rating range [1,10]
	Min         float64 // range [0,5]
	//MinString string  // range [0,5]
	Max float64 // range [0,5]
}

type Language struct {
	Id    ID
	Name  lang.TranslationMap
	Level lang.TranslationMap
}

type LeftSideActionType int

const (
	TechSkill     = LeftSideActionType(1)
	Interests     = LeftSideActionType(2)
	SoftSkills    = LeftSideActionType(3)
	Languages     = LeftSideActionType(4)
	Hobbies       = LeftSideActionType(5)
	TimeAmount    = LeftSideActionType(6)
	Quotation     = LeftSideActionType(7)
	PersonalGoals = LeftSideActionType(8)
	MoneyAmount   = LeftSideActionType(9)
)

func (x LeftSideActionType) FirstSide() bool {
	switch x {
	case TechSkill, Interests, SoftSkills, Languages, Hobbies:
		return true
	default:
		return false
	}
}

type LeftSideActionTypes []LeftSideActionType

func (x LeftSideActionTypes) Index(l LeftSideActionType) (int, bool) {
	for i, a := range x {
		if a == l {
			return i, true
		}
	}
	return 0, false
}

type ExperiencePart int

const (
	ExperiencePart_description = ExperiencePart(iota) + 1
	ExperiencePart_jobPosition
	ExperiencePart_project
	ExperiencePart_role
	ExperiencePart_techStack
)

type Experience struct {
	Id               ID
	JobPosition      lang.TranslationMap
	JobPosition2     lang.TranslationMap
	JobPosition3     lang.TranslationMap
	Company          string
	StartTime        string
	EndTime          string
	Description      lang.TranslationMap
	Description2     lang.TranslationMap
	Description3     lang.TranslationMap
	Project          lang.TranslationMap
	Project2         lang.TranslationMap
	Project3         lang.TranslationMap
	Role             lang.TranslationMap
	Role2            lang.TranslationMap
	Role3            lang.TranslationMap
	TechStack        lang.TranslationMap
	TechStack2       lang.TranslationMap
	TechStack3       lang.TranslationMap
	FutureExperience bool
	DocumentLinks    []string
}

type Education struct {
	Id            ID
	Graduation    lang.TranslationMap
	Institute     string
	StartTime     string
	EndTime       string
	Focus         lang.TranslationMap
	FinalGrade    lang.TranslationMap
	DocumentLinks []string
}

type Publication struct {
	Id            ID
	Title         lang.TranslationMap
	Publisher     lang.TranslationMap
	StartTime     string
	EndTime       string
	Focus         lang.TranslationMap
	DocumentLinks []string
}

type Award struct {
	Id            ID
	Title         lang.TranslationMap
	Institute     lang.TranslationMap
	StartTime     string
	EndTime       string
	Focus         lang.TranslationMap
	DocumentLinks []string
}

type ProfileAddress struct {
	Street  string
	Zip     string
	City    lang.TranslationMap
	Country lang.TranslationMap
}

type FunFacts struct {
	EarlyBird_NightOwl int // [0 - 1 - 2]

}

type Profile struct {
	FirstName          string
	LastName           string
	Title              string
	Image              string // path
	Nationality        lang.TranslationMap
	Birthday           string
	Address            ProfileAddress
	Email              string
	Phone              string
	SignPath           string   // path to the sign image
	Websites           []string // url
	Attachment         []string
	LeftSideActionType LeftSideActionTypes // add here the order of leftSideActions

	// professional/tech skills
	CustomProfessionalSkillLabel lang.TranslationMap
	ProfessionalSkills           []LeftSideAction // should contains all professional skills you have

	// soft skills
	CustomSoftSkillLabel lang.TranslationMap
	SoftSkills           []LeftSideAction // should contains all soft skills you have

	// interest
	CustomInterestLabel lang.TranslationMap
	Interest            []LeftSideAction

	// language
	CustomLanguageLabel lang.TranslationMap
	Language            []Language

	// hobbies
	CustomHobbiesLabel lang.TranslationMap
	Hobbies            []LeftSideAction

	// time amount
	CustomTimeAmountLabel lang.TranslationMap
	TimeAmount            []LeftSideAction

	// money amount
	CustomMoneyAmountLabel lang.TranslationMap
	MoneyAmount            []LeftSideAction

	// about me
	CustomAboutMeTextLabel lang.TranslationMap
	GeneralAboutMeText     lang.TranslationMap // your general profile text, this text can contains tex elements

	// motivation
	CustomMyMotivationTextLabel lang.TranslationMap
	GeneralMyMotivationText     lang.TranslationMap // your general motivation text, this text can contains tex elements

	// values
	CustomValuesLabel lang.TranslationMap
	GeneralValuesText lang.TranslationMap // your values text, this text can contains tex elements

	// main question
	CustomMainQuestionTextLabel lang.TranslationMap
	GeneralMainQuestionText     lang.TranslationMap // your main question text, this text can contains tex elements

	// experience
	CustomExperienceLabel lang.TranslationMap
	Experience            []Experience

	// education
	CustomEducationLabel lang.TranslationMap
	Education            []Education

	// publication
	CustomPublicationLabel lang.TranslationMap
	Publication            []Publication

	// award
	CustomAwardLabel lang.TranslationMap
	Award            []Award
}
