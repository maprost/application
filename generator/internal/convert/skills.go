package convert

import (
	"fmt"

	"github.com/maprost/application/generator/genmodel"
)

func setCVSkills(prf *genmodel.Profile) (err error) {
	skillSize := len(prf.Company.Skills)

	// no skills set in company -> use first of cv config
	if skillSize == 0 {
		skillSize = maxSkillSize(len(prf.CV.Skills))

		// use the first #skillSize from the CV list
		prf.CV.Skills = prf.CV.Skills[:skillSize]
		return
	}

	skillMap := createSkillMap(prf.CV.Skills)
	skillSize = maxSkillSize(skillSize)
	skills := make([]genmodel.Skill, skillSize)
	for i, skillID := range prf.Company.Skills {
		if i == skillSize {
			break
		}

		skill, ok := skillMap[skillID]
		if !ok {
			err = fmt.Errorf("SkillID %d is not known", skillID)
			return
		}
		skills[i] = skill
	}

	prf.CV.Skills = skills
	return
}

func maxSkillSize(size int) int {
	if size < maxSkills {
		return size
	}
	return maxSkills
}

func createSkillMap(skills []genmodel.Skill) (result map[genmodel.SkillID]genmodel.Skill) {
	result = make(map[genmodel.SkillID]genmodel.Skill)
	for _, skill := range skills {
		result[skill.ID] = skill
	}
	return
}
