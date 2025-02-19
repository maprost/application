package texmodel

type RSA struct {
	Label   string
	TexList []string

	HasDocumentLinks bool //
	HasExternalLinks bool //
	HasProjects      bool
	HasRole          bool
	HasTechStack     bool
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
	Time                string
	QuitReason          string
	Focus               string
	DocumentLinks       []string
	ExternalLinks       []string
}

type Publication struct {
	Title           string
	Publisher       string
	Time            string
	Content         string
	ContentShortLsa string
	DocumentLinks   []string
	ExternalLinks   []string
}

type Award struct {
	Title           string
	Institute       string
	Time            string
	Content         string
	ContentShortLsa string
	DocumentLinks   []string
	ExternalLinks   []string
}
