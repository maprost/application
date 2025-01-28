package twoside

import (
	"fmt"
	"log"
	"sort"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

type texmodelLeftSideAction struct {
	texmodel.LeftSideAction
	action   genmodel.LeftSideActionType
	orderIdx int
}

func lsaFirstSide(x genmodel.LeftSideActionType) bool {
	switch x {
	case genmodel.TechSkill, genmodel.Interests, genmodel.SoftSkills, genmodel.Languages:
		return true
	default:
		return false
	}
}

func addLsa(data *texmodel.Index, app *genmodel.Application, lang lang.Language) {
	if len(app.Profile.LeftSideActionType) == 0 {
		app.Profile.LeftSideActionType = genmodel.LeftSideActionTypes{
			genmodel.TechSkill,
			genmodel.Interests,
			genmodel.SoftSkills,
			genmodel.Languages,
			genmodel.TimeAmount,
			genmodel.MoneyAmount,
			genmodel.Hobbies,
		}
	}

	var sideOneLsa []texmodelLeftSideAction
	var sideTwoLsa []texmodelLeftSideAction
	addLeftSideAction := func(skill texmodel.LeftSideAction, action genmodel.LeftSideActionType, ok bool, err error) {
		if err != nil {
			fmt.Println("lsa err: " + err.Error())
			return
		}
		if !ok {
			fmt.Println("lsa not ok: " + action.String())
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
			if lsaFirstSide(action) {
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
	addLeftSideAction(convertTimeAmount(app, lang))
	addLeftSideAction(convertMoneyAmount(app, lang))

	conv := func(s []texmodelLeftSideAction) string {
		sort.Slice(s, func(i, j int) bool {
			return s[i].orderIdx < s[j].orderIdx
		})

		var res []texmodel.LeftSideAction
		for _, s := range s {
			res = append(res, s.LeftSideAction)
		}

		fmt.Printf("%+v\n", res)
		resTex, err := compiler.CompileSubTex(templatePath(), "lsa.tex", res)
		if err != nil {
			return "error: " + err.Error()
		}
		return resTex
	}
	data.SideOneLSA = conv(sideOneLsa)
	data.SideTwoLSA = conv(sideTwoLsa)
}

func convertProfSkills(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	res.Type = 1
	res.Label = app.Profile.CustomProfessionalSkillLabel[lang]
	if res.Label == "" {
		res.Label = lang.TechSkills()
	}

	skills, action, err := util.CalculateProfessionalSkills(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	if err != nil {
		return res, action, false, err
	}

	maxSkills := app.JobPosition.TwoSideStyle.ViewProfessionalSkillRatingSize
	if maxSkills == 0 {
		maxSkills = app.Profile.ViewProfessionalSkillRatingSize
	}
	if maxSkills == 0 {
		maxSkills = 5
	}
	if maxSkills+1 == len(skills) {
		maxSkills++
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
	list, action, err := util.CalculateSoftSkills(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	customLabel := app.Profile.CustomSoftSkillLabel[lang]
	defaultLabel := lang.SoftSkills()
	return convertListSkill(list, action, customLabel, defaultLabel, err, lang)
}

func convertInterests(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	list, action, err := util.CalculateInterest(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	customLabel := app.Profile.CustomInterestLabel[lang]
	defaultLabel := lang.Interests()
	return convertListSkill(list, action, customLabel, defaultLabel, err, lang)
}

func convertHobbies(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	list, action, err := util.CalculateHobbies(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	customLabel := app.Profile.CustomHobbiesLabel[lang]
	defaultLabel := lang.Hobbies()
	return convertListSkill(list, action, customLabel, defaultLabel, err, lang)
}

func convertTimeAmount(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {

	log.Printf("convertTimeAmount start\n")

	var res texmodel.LeftSideAction
	res.Type = 1
	res.Label = app.Profile.CustomTimeAmountLabel[lang]
	if res.Label == "" {
		res.Label = lang.TimeAmount()
	}

	log.Printf("convertTimeAmount 2\n")

	skills, action, err := util.CalculateTimeAmount(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	if err != nil {
		return res, action, false, err
	}

	log.Printf("convertTimeAmount 3\n")

	//maxSkills := maxProfessionalSkills
	//if maxProfessionalSkills+1 == len(skills) {
	//	maxSkills = maxProfessionalSkills + 1
	//}

	fulltime := 40.0

	for _, skill := range skills {

		delta := (skill.Max - skill.Min) / fulltime
		deltafull := 1 - (skill.Max / fulltime)

		res.Range = append(res.Range, texmodel.RangeLsa{
			Name:         lang.String(skill.Name),
			Min:          skill.Min / fulltime,
			MinString:    fmt.Sprintf("%.3f", skill.Min/fulltime),
			MinLabel:     fmt.Sprintf("%.0fh\n", skill.Min),
			Max:          skill.Max / fulltime,
			MaxString:    fmt.Sprintf("%.3f", skill.Max/fulltime),
			MaxLabel:     fmt.Sprintf("%.0fh\n", skill.Max),
			DeltaMaxMin:  fmt.Sprintf("%f", delta),
			DeltaMaxFull: fmt.Sprintf("%.3f", deltafull),
		})
	}
	log.Printf("convertTimeAmount 4\n")
	log.Printf("%+v\n", res)
	log.Printf("convertTimeAmount 5\n")

	return res, action, len(skills) > 0, nil
}

func convertMoneyAmount(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {

	log.Printf("convertMoneyAmount start\n")

	var res texmodel.LeftSideAction
	res.Type = 1
	res.Label = app.Profile.CustomMoneyAmountLabel[lang]
	if res.Label == "" {
		res.Label = lang.MoneyAmount()
	}

	log.Printf("convertMoneyAmount 2\n")

	skills, action, err := util.CalculateMoneyAmount(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	if err != nil {
		return res, action, false, err
	}

	log.Printf("convertMoneyAmount 3\n")

	//maxSkills := maxProfessionalSkills
	//if maxProfessionalSkills+1 == len(skills) {
	//	maxSkills = maxProfessionalSkills + 1
	//}

	fullmoney := 100.0
	currencyText := "T€"

	for _, skill := range skills {

		if skill.Max != 0 {
			delta := (skill.Max - skill.Min) / fullmoney
			deltafull := 1 - (skill.Max / fullmoney)

			mainlabel := lang.String(skill.SingleLabel)
			maxlabel := fmt.Sprintf("%.0f%s\n", skill.Max, currencyText)
			if lang.String(skill.SingleLabel) == "" {
				mainlabel = fmt.Sprintf("%.0f%s\n", skill.Min, currencyText)
			} else {
				maxlabel = ""
			}

			res.Range = append(res.Range, texmodel.RangeLsa{
				Name:            lang.String(skill.Name),
				Min:             skill.Min / fullmoney,
				MinString:       fmt.Sprintf("%.3f", skill.Min/fullmoney),
				MinLabel:        mainlabel,
				Max:             skill.Max / fullmoney,
				MaxString:       fmt.Sprintf("%.3f", skill.Max/fullmoney),
				MaxLabel:        maxlabel,
				DeltaMaxMin:     fmt.Sprintf("%f", delta),
				DeltaMaxMinHalf: fmt.Sprintf("%f", delta/2),
				DeltaMaxFull:    fmt.Sprintf("%.3f", deltafull),
			})

			//} else {
			//	res.Range = append(res.Range, texmodel.RangeLsa{
			//		Name:         lang.String(skill.Name),
			//		Min:          skill.Min / fullmoney,
			//		MinString:    fmt.Sprintf("%s", skill.Min),
			//		MinLabel:     fmt.Sprintf("%s", skill.Min),
			//		Max:          skill.Max / fullmoney,
			//		MaxString:    fmt.Sprintf(""),
			//		MaxLabel:     fmt.Sprintf(""),
			//		DeltaMaxMin:  fmt.Sprintf(""),
			//		DeltaMaxFull: fmt.Sprintf(""),
			//	})
		}
	}
	log.Printf("convertMoneyAmount 4\n")
	log.Printf("%+v\n", res)
	log.Printf("convertMoneyAmount 5\n")

	return res, action, len(skills) > 0, nil
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

func convertLanguage(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	langs, action, err := util.CalculateLanguage(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)

	var res texmodel.LeftSideAction
	res.Type = 3
	res.Label = app.Profile.CustomLanguageLabel[lang]
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
