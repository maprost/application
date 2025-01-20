package util

import (
	"fmt"
	"sort"

	"github.com/maprost/application/generator/genmodel"
)

func CalculateProfessionalSkills(app *genmodel.Application, needed map[genmodel.LeftSideActionType][]genmodel.ID) ([]genmodel.LeftSideAction, genmodel.LeftSideActionType, error) {
	return CalculateLsa(app.Profile.ProfessionalSkills, needed, genmodel.TechSkill)
}

func CalculateSoftSkills(app *genmodel.Application, needed map[genmodel.LeftSideActionType][]genmodel.ID) ([]genmodel.LeftSideAction, genmodel.LeftSideActionType, error) {
	return CalculateLsa(app.Profile.SoftSkills, needed, genmodel.SoftSkills)
}

func CalculateInterest(app *genmodel.Application, needed map[genmodel.LeftSideActionType][]genmodel.ID) ([]genmodel.LeftSideAction, genmodel.LeftSideActionType, error) {
	return CalculateLsa(app.Profile.Interest, needed, genmodel.Interests)
}

func CalculateHobbies(app *genmodel.Application, needed map[genmodel.LeftSideActionType][]genmodel.ID) ([]genmodel.LeftSideAction, genmodel.LeftSideActionType, error) {
	return CalculateLsa(app.Profile.Hobbies, needed, genmodel.Hobbies)
}

func CalculateLsa(lsaList []genmodel.LeftSideAction, needed map[genmodel.LeftSideActionType][]genmodel.ID, action genmodel.LeftSideActionType) ([]genmodel.LeftSideAction, genmodel.LeftSideActionType, error) {
	var result []genmodel.LeftSideAction
	neededLsa := needed[action]
	if len(neededLsa) == 0 {
		result = sortLsa(lsaList)
		return result, action, nil
	}

	// collect all needed skills out of the map
	for _, lsaId := range neededLsa {
		var lsa genmodel.LeftSideAction
		found := false
		for _, s := range lsaList {
			if s.Id == lsaId {
				lsa = s
				found = true
			}
		}
		if !found {
			err := fmt.Errorf("can't find lsa Id '%v' in the map'%v'", lsaId, lsaList)
			return nil, action, err
		}
		result = append(result, lsa)
	}
	return result, action, nil
}

// return the map values sorted by ranking
func sortLsa(skills []genmodel.LeftSideAction) []genmodel.LeftSideAction {
	sort.Slice(skills, func(i, j int) bool {
		return skills[i].Rating > skills[j].Rating
	})
	return skills
}

func CalculateLanguage(app *genmodel.Application, needed map[genmodel.LeftSideActionType][]genmodel.ID) ([]genmodel.Language, genmodel.LeftSideActionType, error) {
	action := genmodel.Languages
	neededLanguages := needed[action]
	if len(neededLanguages) == 0 {
		return app.Profile.Language, action, nil
	}

	// collect all needed skills out of the map
	var result []genmodel.Language
	for _, langId := range neededLanguages {
		var lang genmodel.Language
		found := false
		for _, l := range app.Profile.Language {
			if l.Id == langId {
				lang = l
				found = true
			}
		}
		if !found {
			err := fmt.Errorf("can't find lang Id '%v' in the list'%v'", langId, app.Profile.Language)
			return nil, action, err
		}
		result = append(result, lang)
	}
	return result, action, nil
}
