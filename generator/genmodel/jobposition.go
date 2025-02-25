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
	JobPositionFirstLine string
	JobPosition          string
	StartTime            string
	EndTime              string
	Description          string
	TechStack            string
}

func (e FutureExperience) Empty() bool {
	return len(e.JobPosition) == 0
}

type JobPosition struct {
	MainColor   string // please use the HTML color signature: 800000, if the field is empty: standard color will used
	SideColor   string
	ShadowColor string
	ScaleLineBG string
	Color1      string
	Color2      string
	Color3      string
	Color4      string
	Color5      string
	Lang        lang.Language
	Company     string
	Address     JobAddress
	UseMig      bool
	Title       string // of the job
	FileName    string // default is 'application'
	OutputPath  string // will overwrite build(outputPath)

	// customize your profile
	CoverLetterSubject    string   //  your cover letter subject, this text can contain tex elements
	CoverLetterTxt        string   //  your cover letter text, this text can contain tex elements
	CoverLetterAttachment []string //  your cover letter attachment, this text can contain tex elements
	ProfileLabel          string
	ProfileText           string //  your job specific profile text, this text can contain tex elements
	MyMotivationLabel     string
	MyMotivationText      string //  your job specific motivation text, this text can contain tex elements
	MainQuestionLabel     string
	MainQuestionText      string //  your job specific question text, this text can contain tex elements
	TechSkillsLabel       string
	SoftSkillsLabel       string
	MoneyAmount           []LeftSideAction //  your job specific money amount
	TimeAmount            []LeftSideAction //  your job specific time amount
	FutureExperience      FutureExperience // if set, this experience will be the top of your experience
	FutureColorChange     bool
	Attachment            []string

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
