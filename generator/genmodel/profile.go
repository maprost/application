package genmodel

import "github.com/maprost/application/generator/lang"

type SkillID int

type Skill struct {
	Name   lang.TranslationMap
	Rating int // Rating range [1,10]
}

type LeftSideAction int

const (
	TechSkill  = LeftSideAction(1)
	Interests  = LeftSideAction(2)
	SoftSkills = LeftSideAction(3)
	Languages  = LeftSideAction(4)
	Hobbies    = LeftSideAction(5)
)

type Experience struct {
	JobPosition      lang.TranslationMap
	Company          string
	StartTime        string
	EndTime          string
	Description      lang.TranslationMap
	Project          lang.TranslationMap
	Role             lang.TranslationMap
	TechStack        lang.TranslationMap
	FutureExperience bool
	DocumentLinks    []string
}

type Education struct {
	Graduation    lang.TranslationMap
	Institute     string
	StartTime     string
	EndTime       string
	Focus         lang.TranslationMap
	FinalGrade    lang.TranslationMap
	DocumentLinks []string
}

type Language struct {
	Name  lang.TranslationMap
	Level lang.TranslationMap
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
	FirstName                    string
	LastName                     string
	Title                        string
	Image                        string // path
	Nationality                  lang.TranslationMap
	Birthday                     string
	Address                      ProfileAddress
	Email                        string
	Phone                        string
	Websites                     []string // url
	CustomProfessionalSkillLabel lang.TranslationMap
	ProfessionalSkills           map[SkillID]Skill // should contains all professional skills you have
	CustomSoftSkillLabel         lang.TranslationMap
	SoftSkills                   map[SkillID]Skill // should contains all soft skills you have
	CustomInterestLabel          lang.TranslationMap
	Interest                     []lang.TranslationMap
	CustomHobbiesLabel           lang.TranslationMap
	Hobbies                      []lang.TranslationMap
	SignPath                     string // path to the sign image
	CustomExperienceLabel        lang.TranslationMap
	Experience                   []Experience
	CustomEducationLabel         lang.TranslationMap
	Education                    []Education
	CustomLanguageLabel          lang.TranslationMap
	Language                     []Language
	CustomMotivationTextLabel    lang.TranslationMap
	GeneralMotivationText        lang.TranslationMap // your general motivation text, this text can contains tex elements
	Attachment                   []string
	LeftSideAction               []LeftSideAction
}
