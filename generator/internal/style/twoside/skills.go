package twoside

import (
	"github.com/maprost/application/generator/genmodel"
)

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
