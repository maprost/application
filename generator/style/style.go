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
		oss := app.JobPosition.OneSideStyle
		if len(oss.LeftSideActionTypes) == 0 {
			oss.LeftSideActionTypes = genmodel.AllLeftSideActionTypes
		}

		app.JobPosition.TwoSideStyle = genmodel.TwoSideStyle{
			ActivateCoverLetter:      oss.ActivateCoverLetter,
			CoverLetterOnSeparatePdf: oss.CoverLetterOnSeparatePdf,
			// LSA
			Skills:                          oss.Skills,
			RemoveSkills:                    oss.RemoveSkills,
			SideOneLSATypes:                 oss.LeftSideActionTypes,
			SideTwoLSATypes:                 nil,
			ViewProfessionalSkillRatingSize: oss.ViewProfessionalSkillRatingSize,
			// RSA
			SideOneRSATypes: oss.RightSideActionTypes,
			SideTwoRSATypes: nil,
			SideOneRSAItems: math.MaxInt, // put all on side one
			// experience
			Experience:           oss.Experience,
			RemoveExperience:     oss.RemoveExperience,
			ShowExperienceParts:  oss.ShowExperienceParts,
			ShowExperiencePart:   oss.ShowExperiencePart,
			RemoveExperiencePart: oss.RemoveExperiencePart,
			// education
			Education:       oss.Education,
			RemoveEducation: oss.RemoveEducation,
			// publication
			Publication:       oss.Publication,
			RemovePublication: oss.RemovePublication,
			// award
			Award:       oss.Award,
			RemoveAward: oss.RemoveAward,
			// attachment
			NoAttachmentAndHintsPage: oss.NoAttachmentAndHintsPage,
			NoAttachments:            oss.NoAttachments,
			NoDocumentLinks:          oss.NoLinks,
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
