package genmodel

import (
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
	Skills          map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	LeftSideActions LeftSideActionTypes
	Experience      []ID
	Education       []ID
}

type TwoSideStyle struct {
	Skills                     map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	SideOneLeftSideActionTypes LeftSideActionTypes
	SideTwoLeftSideActionTypes LeftSideActionTypes
	Experience                 []ID
	SideOneExperienceSize      int // rest is on side two
	Education                  []ID
	SideOneEducationSize       int // if SideOneExperienceSize is set,
}

type JobPosition struct {
	MainColor string // please use the HTML color signature: 800000, if the field is empty: standard color will used
	SideColor string
	Lang      lang.Language
	Company   string
	Address   JobAddress
	Title     string // of the job

	// customize your profile
	MotivationText   string           //  your job specific motivation text, this text can contains tex elements
	FutureExperience FutureExperience // if set, this experience will be the top of your experience
	Attachment       []string

	// customize style
	Style        Style
	OneSideStyle OneSideStyle
	TwoSideStyle TwoSideStyle
}
