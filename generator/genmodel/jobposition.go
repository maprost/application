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
	ProfileLabel      string
	ProfileText       string //  your job specific profile text, this text can contains tex elements
	MyMotivationLabel string
	MyMotivationText  string //  your job specific motivation text, this text can contains tex elements
	MainQuestionLabel string
	MainQuestionText  string           //  your job specific question text, this text can contains tex elements
	MoneyAmount       []LeftSideAction //  your job specific money amount
	TimeAmount        []LeftSideAction //  your job specific time amount
	FutureExperience  FutureExperience // if set, this experience will be the top of your experience
	Attachment        []string

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
