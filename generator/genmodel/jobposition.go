package genmodel

import (
	"path"
	"runtime"
	"strings"

	"github.com/maprost/application/generator/lang"
)

type JobAddress struct {
	Street  string
	Zip     string
	City    string
	Country string
}

type FutureExperience struct {
	JobPosition string
	StartTime   string
	Description string
	TechStack   string
}

func (e FutureExperience) Empty() bool {
	return len(e.JobPosition) == 0
}

type Style interface {
	Data(application *Application) (data interface{}, err error)
	Files() (path string, mainFile string, subFiles []string)
}

type OneSideStyle struct {
	Skills               map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	RemoveSkills         map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	LeftSideActionTypes  LeftSideActionTypes
	Experience           []ID
	RemoveExperience     []ID
	ShowExperienceParts  int                    // if 0, it's default 1
	ShowExperiencePart   map[ExperiencePart]int // [1,2,3], if not in, it's default 1
	RemoveExperiencePart []ExperiencePart
	Education            []ID
	RemoveEducation      []ID
}

type TwoSideStyle struct {
	Skills                     map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	RemoveSkills               map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	SideOneLeftSideActionTypes LeftSideActionTypes
	SideTwoLeftSideActionTypes LeftSideActionTypes
	// experience
	Experience            []ID
	RemoveExperience      []ID
	ShowExperienceParts   int                    // if 0, it's default 1
	ShowExperiencePart    map[ExperiencePart]int // [1,2,3], if not in, it's default 1
	RemoveExperiencePart  []ExperiencePart
	SideOneExperienceSize int // rest is on side two
	// education
	Education            []ID
	RemoveEducation      []ID
	SideOneEducationSize int // if SideOneExperienceSize is set, SideOneEducationSize is 0
	// publication
	Publication            []ID
	RemovePublication      []ID
	SideOnePublicationSize int
	// award
	Award            []ID
	RemoveAward      []ID
	SideOneAwardSize int
}

type JobPosition struct {
	MainColor  string // please use the HTML color signature: 800000, if the field is empty: standard color will used
	SideColor  string
	Lang       lang.Language
	Company    string
	Address    JobAddress
	Title      string // of the job
	FileName   string // default is 'application'
	OutputPath string // will overwrite build(outputPath)

	// customize your profile
	ProfileText      string           //  your job specific profile text, this text can contains tex elements
	MyMotivationText string           //  your job specific motivation text, this text can contains tex elements
	MainQuestionText string           //  your job specific question text, this text can contains tex elements
	TimeAmount       []LeftSideAction //  your job specific time amount
	FutureExperience FutureExperience // if set, this experience will be the top of your experience
	Attachment       []string

	// customize style
	Style        Style
	OneSideStyle OneSideStyle
	TwoSideStyle TwoSideStyle
}

func OutputPath() string {
	_, file, _, _ := runtime.Caller(1)
	p := path.Dir(file)
	if !strings.HasSuffix(p, "/") {
		p += "/"
	}
	return p
}
