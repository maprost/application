package texmodel

import "fmt"

type RSA struct {
	Label   string
	TexList []string

	HasDocumentLinks bool //
	HasExternalLinks bool //
	HasProjects      bool
	HasRole          bool
	HasTechStack     bool
}

type RunningLanguage struct {
	Name  string
	Value string
}

func (x RunningLanguage) String() string {
	return fmt.Sprintf("%s: %s", x.Name, x.Value)
}

type RunningLanguages []RunningLanguage

func (x RunningLanguages) String() string {
	var res string
	for _, e := range x {
		if res != "" {
			res += " – "
		}
		res += e.String()
	}
	return res
}

func (x RSA) Split(idx int) (RSA, RSA) {
	first := x // copy
	first.TexList = x.TexList[:idx]
	second := RSA{
		TexList: x.TexList[idx:],
	}
	return first, second
}

type Experience struct {
	Company           string
	Languages         RunningLanguages
	PositionFirstLine string
	Position          string
	Time              string
	TimeSecondLine    string
	QuitReason        string
	Description       string
	Project           string // please add here your pet projects
	Role              string // please add here more roles you did
	Tech              string
	DocumentLinks     []string
	ExternalLinks     []string
	FutureExperience  bool
	FutureColorChange bool
}

type Education struct {
	GraduationFirstLine string
	Graduation          string
	FinalGrade          string
	Institute           string
	Languages           RunningLanguages
	Time                string
	QuitReason          string
	Focus               string
	DocumentLinks       []string
	ExternalLinks       []string
}

type Publication struct {
	Title           string
	Publisher       string
	Languages       RunningLanguages
	Time            string
	Content         string
	ContentShortLsa string
	DocumentLinks   []string
	ExternalLinks   []string
}

type Award struct {
	Title           string
	Institute       string
	Languages       RunningLanguages
	Time            string
	Content         string
	ContentShortLsa string
	DocumentLinks   []string
	ExternalLinks   []string
}
