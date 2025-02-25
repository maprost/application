package style

import (
	"errors"
	"math"

	"github.com/maprost/application/generator/genmodel"
	twoside "github.com/maprost/application/generator/internal/style/twoside"
)

type Style int

const (
	OneSide = Style(iota)
	TwoSide
)

func (s Style) Data(app *genmodel.Application, outputPath string) (data interface{}, err error) {
	switch s {
	case OneSide:
		// convert oneSide into twoSide
		oneSideStyle := app.JobPosition.OneSideStyle
		if len(oneSideStyle.LeftSideActionTypes) == 0 {
			oneSideStyle.LeftSideActionTypes = genmodel.AllLeftSideActionTypes
		}

		app.JobPosition.TwoSideStyle = genmodel.TwoSideStyle{
			// LSA
			Skills:                          oneSideStyle.Skills,
			RemoveSkills:                    oneSideStyle.RemoveSkills,
			SideOneLSATypes:                 oneSideStyle.LeftSideActionTypes,
			SideTwoLSATypes:                 nil,
			ViewProfessionalSkillRatingSize: oneSideStyle.ViewProfessionalSkillRatingSize,
			// RSA
			SideOneRSATypes: oneSideStyle.RightSideActionTypes,
			SideTwoRSATypes: nil,
			SideOneRSAItems: math.MaxInt, // put all on side one
			// experience
			Experience:           oneSideStyle.Experience,
			RemoveExperience:     oneSideStyle.RemoveExperience,
			ShowExperienceParts:  oneSideStyle.ShowExperienceParts,
			ShowExperiencePart:   oneSideStyle.ShowExperiencePart,
			RemoveExperiencePart: oneSideStyle.RemoveExperiencePart,
			// education
			Education:       oneSideStyle.Education,
			RemoveEducation: oneSideStyle.RemoveEducation,
			// publication
			Publication:       oneSideStyle.Publication,
			RemovePublication: oneSideStyle.RemovePublication,
			// award
			Award:       oneSideStyle.Award,
			RemoveAward: oneSideStyle.RemoveAward,
		}
		app.JobPosition.Style = TwoSide
		return twoside.Data(app, outputPath)

	case TwoSide:
		return twoside.Data(app, outputPath)

	}

	err = errors.New("Style not found.")
	return
}

func (s Style) Files() (path string, mainFile string, subFiles []string) {
	switch s {
	case OneSide:
		return twoside.Files()
	case TwoSide:
		return twoside.Files()
	}
	return
}
