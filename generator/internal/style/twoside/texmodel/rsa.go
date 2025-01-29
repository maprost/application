package texmodel

type RSA struct {
	Label   string
	TexList []string

	HasDocumentLinks bool //
	HasProjects      bool
	HasRole          bool
	HasTechStack     bool
}

type Experience struct {
	Company           string
	Position          string
	Time              string
	TimeSecondLine    string
	QuitReason        string
	Description       string
	Project           string // please add here your pet projects
	Role              string // please add here more roles you did
	Tech              string
	DocumentLinks     []string
	FutureExperience  bool
	FutureColorChange bool
}

type Education struct {
	Graduation    string
	FinalGrade    string
	Institute     string
	Time          string
	QuitReason    string
	Focus         string
	DocumentLinks []string
}

type Publication struct {
	Title           string
	Publisher       string
	Time            string
	Content         string
	ContentShortLsa string
	DocumentLinks   []string
}

type Award struct {
	Title           string
	Institute       string
	Time            string
	Content         string
	ContentShortLsa string
	DocumentLinks   []string
}
