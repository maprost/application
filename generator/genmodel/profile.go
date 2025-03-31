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
	Max            float64 // range [0,5]
	Full           float64
	CurrencyEnding string
	MustBeSelected bool // is hidden by default
}

type Language struct {
	Id             ID
	Name           lang.TranslationMap
	Level          lang.TranslationMap
	DocumentLinks  []string
	ShowRefExplain bool
	MustBeSelected bool // is hidden by default
}

type RunningLanguage struct {
	Name     lang.TranslationMap
	Language lang.Language
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
	Id                   ID
	Languages            []RunningLanguage
	JobPositionFirstLine lang.TranslationMap
	JobPosition          lang.TranslationMap
	JobPosition2         lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	JobPosition3         lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	Company              string
	StartTime            string
	EndTime              string
	QuitReason           lang.TranslationMap
	Description          lang.TranslationMap
	Description2         lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	Description3         lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	Project              lang.TranslationMap
	Project2             lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	Project3             lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	Role                 lang.TranslationMap
	Role2                lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	Role3                lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	TechStack            lang.TranslationMap
	TechStack2           lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	TechStack3           lang.TranslationMap // if empty --> will be like 1; if must stay empty: "xxx"
	FutureExperience     bool
	FutureColorChange    bool
	DocumentLinks        []string
	ExternalLinks        []string
	MustBeSelected       bool // is hidden by default
}

type Education struct {
	Id                  ID
	Languages           []RunningLanguage
	GraduationFirstLine lang.TranslationMap
	Graduation          lang.TranslationMap
	Institute           string
	StartTime           string
	EndTime             string
	QuitReason          lang.TranslationMap
	Focus               lang.TranslationMap
	FinalGrade          lang.TranslationMap
	DocumentLinks       []string
	ExternalLinks       []string
	MustBeSelected      bool // is hidden by default
}

type Publication struct {
	Id              ID
	Languages       []RunningLanguage
	Title           lang.TranslationMap
	SubTitle        lang.TranslationMap
	Publisher       lang.TranslationMap
	Date            string
	CoverImage      string
	Content         lang.TranslationMap
	ContentShortLsa lang.TranslationMap
	DocumentLinks   []string
	ExternalLinks   []string
	MustBeSelected  bool // is hidden by default
}

type Award struct {
	Id              ID
	Languages       []RunningLanguage
	Title           lang.TranslationMap
	Institute       lang.TranslationMap
	Date            string
	Content         lang.TranslationMap
	ContentShortLsa lang.TranslationMap
	DocumentLinks   []string
	ExternalLinks   []string
	MustBeSelected  bool // is hidden by default
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
	FirstName           string
	LastName            string
	Title               lang.TranslationMap
	Image               string // path
	Nationality         lang.TranslationMap
	NationalityMig      lang.TranslationMap
	Birthday            string
	Address             ProfileAddress // will be used inside the cl/cv
	Email               string
	Phone               string
	SignPath            string   // path to the sign image
	Websites            []string // url
	Attachment          []string
	LeftSideActionType  LeftSideActionTypes  // add here the order of leftSideActions
	RightSideActionType RightSideActionTypes // add here the order of leftSideActions

	// cover letter
	CoverLetterSubject    lang.TranslationMap
	CoverLetterTxtIntro   lang.TranslationMap
	CoverLetterTxt        lang.TranslationMap
	CoverLetterTxtOutro   lang.TranslationMap
	CoverLetterAttachment []lang.TranslationMap

	// professional/tech skills
	CustomProfessionalSkillLabel    lang.TranslationMap
	ProfessionalSkills              []LeftSideAction // should contain all professional skills you have
	ViewProfessionalSkillRatingSize int

	// soft skills
	CustomSoftSkillLabel lang.TranslationMap
	SoftSkills           []LeftSideAction // should contain all soft skills you have

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

	// profile
	CustomProfilTextLabel lang.TranslationMap
	GeneralProfileText    lang.TranslationMap // your general profile text, this text can contain tex elements

	// motivation
	CustomMyMotivationTextLabel lang.TranslationMap
	GeneralMyMotivationText     lang.TranslationMap // your general motivation text, this text can contain tex elements

	// values
	CustomValuesLabel lang.TranslationMap
	GeneralValuesText lang.TranslationMap // your values text, this text can contain tex elements

	// main question
	CustomMainQuestionTextLabel lang.TranslationMap
	GeneralMainQuestionText     lang.TranslationMap // your main question text, this text can contain tex elements

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
