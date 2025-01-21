package style

import (
	"errors"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside"
)

type Style int

const (
	OneSide = Style(iota)
	TwoSide
)

func (s Style) Data(application *genmodel.Application) (data interface{}, err error) {
	switch s {
	case OneSide:
		// convert oneSide into twoSide
		oneSideStyle := application.JobPosition.OneSideStyle
		application.JobPosition.TwoSideStyle = genmodel.TwoSideStyle{
			Skills:                     oneSideStyle.Skills,
			RemoveSkills:               oneSideStyle.RemoveSkills,
			SideOneLeftSideActionTypes: oneSideStyle.LeftSideActionTypes,
			SideTwoLeftSideActionTypes: nil,
			Experience:                 oneSideStyle.Experience,
			RemoveExperience:           oneSideStyle.RemoveExperience,
			ShowExperiencePart:         oneSideStyle.ShowExperiencePart,
			ShowExperienceParts:        oneSideStyle.ShowExperienceParts,
			RemoveExperiencePart:       oneSideStyle.RemoveExperiencePart,
			SideOneExperienceSize:      len(application.Profile.Experience),
			Education:                  oneSideStyle.Education,
			RemoveEducation:            oneSideStyle.RemoveEducation,
			SideOneEducationSize:       len(application.Profile.Education),
		}
		application.JobPosition.Style = TwoSide
		return twoside.Data(application)

	case TwoSide:
		return twoside.Data(application)
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
