package twoside

import (
	"sort"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

type texmodelLeftSideAction struct {
	texmodel.LeftSideAction
	action   genmodel.LeftSideActionType
	orderIdx int
}

func addLsa(data *texmodel.Index, app *genmodel.Application, lang lang.Language) {
	if len(app.Profile.LeftSideActionType) == 0 {
		app.Profile.LeftSideActionType = genmodel.LeftSideActionTypes{
			genmodel.TechSkill,
			genmodel.Interests,
			genmodel.SoftSkills,
			genmodel.Languages,
			genmodel.Hobbies,
		}
	}

	var sideOneLsa []texmodelLeftSideAction
	var sideTwoLsa []texmodelLeftSideAction
	addLeftSideAction := func(skill texmodel.LeftSideAction, action genmodel.LeftSideActionType, ok bool, err error) {
		if err != nil {
			return
		}
		if !ok {
			return
		}

		if idx, ok := app.JobPosition.TwoSideStyle.SideOneLeftSideActionTypes.Index(action); ok {
			sideOneLsa = append(sideOneLsa, texmodelLeftSideAction{
				LeftSideAction: skill,
				action:         action,
				orderIdx:       idx,
			})
		} else if idx, ok := app.JobPosition.TwoSideStyle.SideTwoLeftSideActionTypes.Index(action); ok {
			sideTwoLsa = append(sideTwoLsa, texmodelLeftSideAction{
				LeftSideAction: skill,
				action:         action,
				orderIdx:       idx,
			})
		} else if idx, ok := app.Profile.LeftSideActionType.Index(action); ok {
			if action.FirstSide() {
				sideOneLsa = append(sideOneLsa, texmodelLeftSideAction{
					LeftSideAction: skill,
					action:         action,
					orderIdx:       idx,
				})
			} else {
				sideTwoLsa = append(sideTwoLsa, texmodelLeftSideAction{
					LeftSideAction: skill,
					action:         action,
					orderIdx:       idx,
				})
			}
		}
	}
	addLeftSideAction(convertProfSkills(app, lang))
	addLeftSideAction(convertSoftSkills(app, lang))
	addLeftSideAction(convertHobbies(app, lang))
	addLeftSideAction(convertInterests(app, lang))
	addLeftSideAction(convertLanguage(app, lang))

	conv := func(s []texmodelLeftSideAction) []texmodel.LeftSideAction {
		sort.Slice(s, func(i, j int) bool {
			return s[i].orderIdx < s[j].orderIdx
		})

		var res []texmodel.LeftSideAction
		for _, s := range s {
			res = append(res, s.LeftSideAction)
		}
		return res
	}
	data.SideOneLeftSideAction = conv(sideOneLsa)
	data.SideTwoLeftSideAction = conv(sideTwoLsa)
}

func convertProfSkills(application *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	res.Type = 1
	res.Label = application.Profile.CustomProfessionalSkillLabel[lang]
	if res.Label == "" {
		res.Label = lang.TechSkills()
	}

	skills, action, err := util.CalculateProfessionalSkills(application, application.JobPosition.TwoSideStyle.Skills)
	if err != nil {
		return res, action, false, err
	}

	maxSkills := maxProfessionalSkills
	if maxProfessionalSkills+1 == len(skills) {
		maxSkills = maxProfessionalSkills + 1
	}

	for i, skill := range skills {
		if i < maxSkills {
			res.Ratings = append(res.Ratings, texmodel.RatingLsa{
				Name:   lang.String(skill.Name),
				Rating: skill.Rating,
			})
		} else {
			if i > maxSkills {
				res.OtherRatings += ", "
			}
			res.OtherRatings += lang.String(skill.Name)
		}
	}
	return res, action, len(skills) > 0, nil
}

func convertSoftSkills(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	skills, action, err := util.CalculateSoftSkills(app, app.JobPosition.TwoSideStyle.Skills)
	customLabel := app.Profile.CustomSoftSkillLabel[lang]
	defaultLabel := lang.SoftSkills()
	return convertListSkill(skills, action, customLabel, defaultLabel, err, lang)
}

func convertInterests(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	skills, action, err := util.CalculateInterest(app, app.JobPosition.TwoSideStyle.Skills)
	customLabel := app.Profile.CustomInterestLabel[lang]
	defaultLabel := lang.Interests()
	return convertListSkill(skills, action, customLabel, defaultLabel, err, lang)
}

func convertHobbies(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	skills, action, err := util.CalculateHobbies(app, app.JobPosition.TwoSideStyle.Skills)
	customLabel := app.Profile.CustomHobbiesLabel[lang]
	defaultLabel := lang.Hobbies()
	return convertListSkill(skills, action, customLabel, defaultLabel, err, lang)
}

func convertListSkill(skills []genmodel.LeftSideAction, action genmodel.LeftSideActionType, customLabel string, defaultLabel string, err error, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	res.Type = 2

	res.Label = customLabel
	if res.Label == "" {
		res.Label = defaultLabel
	}

	if err != nil {
		return res, action, false, err
	}

	for i, skill := range skills {
		if i > 0 {
			res.List += ", "
		}
		res.List += lang.String(skill.Name)
	}
	return res, action, len(skills) > 0, nil
}

func convertLanguage(application *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	langs, action, err := util.CalculateLanguage(application, application.JobPosition.TwoSideStyle.Skills)

	var res texmodel.LeftSideAction
	res.Type = 3
	res.Label = application.Profile.CustomLanguageLabel[lang]
	if res.Label == "" {
		res.Label = lang.Languages()
	}
	if err != nil {
		return res, action, false, err
	}

	for _, l := range langs {
		res.Languages = append(res.Languages, texmodel.Language{
			Name:  lang.String(l.Name),
			Level: lang.String(l.Level),
		})
	}
	return res, action, len(langs) > 0, nil
}
