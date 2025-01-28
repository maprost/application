package genmodel

type Style interface {
	Data(application *Application) (data interface{}, err error)
	Files() (path string, mainFile string, subFiles []string)
}

type OneSideStyle struct {
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
}

type TwoSideStyle struct {
	// LSA
	Skills                          map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	RemoveSkills                    map[LeftSideActionType][]ID // if nothing is selected, it will use everything from profile
	SideOneLSATypes                 LeftSideActionTypes
	SideTwoLSATypes                 LeftSideActionTypes
	ViewProfessionalSkillRatingSize int

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
}
