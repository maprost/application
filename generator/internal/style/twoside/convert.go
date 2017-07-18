package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
)

func initIndex(application *genmodel.Application) (texmodel.Index, error) {
	return texmodel.Index{
		MainColor: application.JobPosition.MainColor,
	}, nil
}

func createFirstPageData(application *genmodel.Application) (texmodel.FirstPage, error) {
	return texmodel.FirstPage{
		Name:  application.Profile.FirstName + " " + application.Profile.LastName,
		Title: application.JobPosition.Title,
		Image: application.Profile.Image,
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

func createCVData(application *genmodel.Application) (texmodel.CV, error) {
	var err error

	return texmodel.CV{
		Name:  application.Profile.FirstName + " " + application.Profile.LastName,
		Title: application.JobPosition.Title,
		Image: application.Profile.Image,
		Email: application.Profile.Email,
	}, err
}

func setCVSkills(prf *genmodel.Profile) (err error) {
	//skillSize := len(prf.Company.ProfSkills)

	//// no skills set in jobposition -> use first of cv config
	//if skillSize == 0 {
	//	skillSize = maxSkillSize(len(prf.CV.ProfSkills))
	//
	//	// use the first #skillSize from the CV list
	//	prf.CV.ProfSkills = prf.CV.ProfSkills[:skillSize]
	//	return
	//}
	//
	//skillMap := createSkillMap(prf.CV.ProfSkills)
	//skillSize = maxSkillSize(skillSize)
	//skills := make([]genmodel.Skill, skillSize)
	//for i, skillID := range prf.Company.ProfSkills {
	//	if i == skillSize {
	//		break
	//	}
	//
	//	skill, ok := skillMap[skillID]
	//	if !ok {
	//		err = fmt.Errorf("SkillID %d is not known", skillID)
	//		return
	//	}
	//	skills[i] = skill
	//}
	//
	//prf.CV.ProfSkills = skills
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
