package genmodel

type Style interface {
	Data(app *Application, outputPath string) (data interface{}, err error)
	Files() (path string, mainFile string, subFiles []string)
}

type OneSideStyle struct {
	// cover letter
	ActivateCoverLetter      bool
	CoverLetterOnSeparatePdf bool

	// LSA
	Skills                          map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	ViewSkillRatingSize             map[LeftSideActionType]int
	RemoveSkills                    map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	LeftSideActionTypes             LeftSideActionTypes
	ViewProfessionalSkillRatingSize int

	// RSA
	RightSideActionTypes RightSideActionTypes

	// experience
	Experience           []ID
	RemoveExperience     []ID
	ShowExperienceParts  int                    // if 0, it's default 1
	ShowExperiencePart   map[ExperiencePart]int // [1,2,3], if not in, it's default 1
	RemoveExperiencePart []ExperiencePart

	// education
	Education       []ID
	RemoveEducation []ID

	// publication
	Publication       []ID
	RemovePublication []ID

	// award
	Award       []ID
	RemoveAward []ID

	// attachment
	NoAttachmentAndHintsPage bool
	NoAttachments            bool
	NoLinks                  bool
}

type TwoSideStyle struct {
	// cover letter
	ActivateCoverLetter      bool
	CoverLetterOnSeparatePdf bool

	// LSA
	Skills                          map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	RemoveSkills                    map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	ViewProfessionalSkillRatingSize int
	SideOneLSATypes                 LeftSideActionTypes
	SideTwoLSATypes                 LeftSideActionTypes

	// RSA
	SideOneRSATypes RightSideActionTypes
	SideTwoRSATypes RightSideActionTypes
	SideOneRSAItems int // motivation(count 1), mainQuest(count 1), experience (count as len(exp)), education .....

	// experience
	Experience           []ID
	RemoveExperience     []ID
	ShowExperienceParts  int                    // if 0, it's default 1
	ShowExperiencePart   map[ExperiencePart]int // [1,2,3], if not in, it's default 1
	RemoveExperiencePart []ExperiencePart

	// education
	Education       []ID
	RemoveEducation []ID

	// publication
	Publication       []ID
	RemovePublication []ID

	// award
	Award       []ID
	RemoveAward []ID

	// attachment
	NoAttachmentAndHintsPage bool
	NoAttachments            bool
	NoDocumentLinks          bool
}
