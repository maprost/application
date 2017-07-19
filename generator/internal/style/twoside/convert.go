package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
)

func initIndex(application *genmodel.Application) (texmodel.Index, error) {
	return texmodel.Index{
		MainColor: util.DefaultColor(application.JobPosition.MainColor),
	}, nil
}

func createFirstPageData(application *genmodel.Application) (texmodel.FirstPage, error) {
	return texmodel.FirstPage{
		Name:  application.Profile.FirstName + " " + application.Profile.LastName,
		Title: application.JobPosition.Title,
		Image: util.DefaultImage(application.Profile.Image),
	}, nil
}

func createCoverLetterData(application *genmodel.Application) (texmodel.CoverLetter, error) {
	return texmodel.CoverLetter{
		Name:          application.Profile.FirstName + " " + application.Profile.LastName,
		Street:        application.Profile.Address.Street,
		Zip:           application.Profile.Address.Zip + " " + application.Profile.Address.City,
		CompanyName:   application.JobPosition.Company,
		CompanyStreet: application.JobPosition.Address.Street,
		CompanyZip:    application.JobPosition.Address.Zip + " " + application.JobPosition.Address.City,
		Text:          application.JobPosition.LetterText,
	}, nil
}

func createCVData(application *genmodel.Application) (data texmodel.CV, err error) {
	profSkills, otherProfSkills, err := convertProfSkills(application)
	if err != nil {
		return
	}

	softSkills, err := convertSoftSkills(application)
	if err != nil {
		return
	}

	data = texmodel.CV{
		Name:            application.Profile.FirstName + " " + application.Profile.LastName,
		Title:           application.JobPosition.Title,
		Image:           util.DefaultImage(application.Profile.Image),
		Email:           application.Profile.Email,
		OtherProfSkills: otherProfSkills,
		ProfSkills:      profSkills,
		SoftSkills:      softSkills,
	}
	return
}

func convertProfSkills(application *genmodel.Application) (profSkills []texmodel.Skill, otherProfSkills string, err error) {
	skills, err := util.CalculatProfessionalSkills(application)
	if err != nil {
		return
	}

	for i, skill := range skills {
		if i < maxProfessionalSkills {
			profSkills = append(profSkills, texmodel.Skill{
				Name:   skill.Name,
				Rating: skill.Rating,
			})
		} else {
			if i >= maxProfessionalSkills {
				otherProfSkills += " ,"
			}
			otherProfSkills += skill.Name
		}
	}
	return
}

func convertSoftSkills(application *genmodel.Application) (softSkills string, err error) {
	skills, err := util.CalculateSoftSkills(application)
	if err != nil {
		return
	}

	for i, skill := range skills {
		if i > 0 {
			softSkills += " ,"
		}
		softSkills += skill.Name
	}
	return
}

//func maxSkillSize(size int) int {
//	if size < maxSkills {
//		return size
//	}
//	return maxSkills
//}
//
//func createSkillMap(skills []genmodel.Skill) (result map[genmodel.SkillID]genmodel.Skill) {
//	result = make(map[genmodel.SkillID]genmodel.Skill)
//	for _, skill := range skills {
//		result[skill.ID] = skill
//	}
//	return
//}
