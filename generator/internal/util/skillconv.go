package util

import (
	"fmt"
	"github.com/maprost/application/generator/genmodel"
	"sort"
)

func CalculateProfessionalSkills(application *genmodel.Application) ([]genmodel.Skill, error) {
	return calculateSkills(application.Profile.ProfessionalSkills, application.JobPosition.ProfessionalSkills)
}

func CalculateSoftSkills(application *genmodel.Application) ([]genmodel.Skill, error) {
	return calculateSkills(application.Profile.SoftSkills, application.JobPosition.SoftSkills)
}

func calculateSkills(skillMap map[genmodel.SkillID]genmodel.Skill, neededSkills []genmodel.SkillID) (result []genmodel.Skill, err error) {
	if len(neededSkills) == 0 {
		result = convertMapToList(skillMap)
		return
	}

	// collect all needed skills out of the map
	for _, skillID := range neededSkills {
		skill, ok := skillMap[skillID]
		if !ok {
			err = fmt.Errorf("Can't find skill ID '%v' in the map'%v'", skillID, skillMap)
			return
		}
		result = append(result, skill)
	}
	return
}

// return the map values sorted by ranking
func convertMapToList(skillMap map[genmodel.SkillID]genmodel.Skill) (result []genmodel.Skill) {
	for _, skill := range skillMap {
		result = append(result, skill)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Rating > result[j].Rating
	})
	return
}
