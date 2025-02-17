package twoside

import (
	"fmt"
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

func addLsa(data *texmodel.Index, app *genmodel.Application, local lang.Language) {
	style := app.JobPosition.TwoSideStyle
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
	addLeftSideAction := func(fn func(app *genmodel.Application, local lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error)) {
		skill, action, ok, err := fn(app, local)
		if err != nil {
			fmt.Println("lsa err: " + err.Error())
			return
		}
		if !ok {
			fmt.Println("lsa not ok: " + action.String())
			return
		}

		if len(style.SideOneLSATypes) > 0 || len(style.SideTwoLSATypes) > 0 {
			if idx, ok := style.SideOneLSATypes.Index(action); ok {
				sideOneLsa = append(sideOneLsa, texmodelLeftSideAction{
					LeftSideAction: skill,
					action:         action,
					orderIdx:       idx,
				})
			} else if idx, ok := style.SideTwoLSATypes.Index(action); ok {
				sideTwoLsa = append(sideTwoLsa, texmodelLeftSideAction{
					LeftSideAction: skill,
					action:         action,
					orderIdx:       idx,
				})
			} else {
				// do nothing
			}
		} else {
			if idx, ok := app.Profile.LeftSideActionType.Index(action); ok {
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
	}

	addLeftSideAction(lsaConvProfSkills)
	addLeftSideAction(lsaConvSoftSkills)
	addLeftSideAction(lsaConvHobbies)
	addLeftSideAction(lsaConvInterests)
	addLeftSideAction(lsaConvLanguage)
	addLeftSideAction(lsaConvTimeAmount)
	addLeftSideAction(lsaConvMoneyAmount)
	addLeftSideAction(lsaConvPublication)
	addLeftSideAction(lsaConvAward)

	conv := func(s []texmodelLeftSideAction) string {
		sort.Slice(s, func(i, j int) bool {
			return s[i].orderIdx < s[j].orderIdx
		})

		var res []texmodel.LeftSideAction
		for _, s := range s {
			res = append(res, s.LeftSideAction)
		}

		fmt.Printf("%+v\n", res)
		resTex, err := compiler.CompileSubTex(templatePath(), "lsa.tex", texmodel.LSAIndex{
			Label: local,
			List:  res,
		})
		if err != nil {
			return "error: " + err.Error()
		}
		return resTex
	}
	data.SideOneLSA = conv(sideOneLsa)
	data.SideTwoLSA = conv(sideTwoLsa)
}

func lsaConvProfSkills(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	res.Label = app.Profile.CustomProfessionalSkillLabel[lang]
	if res.Label == "" {
		res.Label = lang.TechSkills()
	}
	// TODO improve
	if app.JobPosition.TechSkillsLabel != "" {
		res.Label = app.JobPosition.TechSkillsLabel
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

func lsaConvSoftSkills(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	list, action, err := util.CalculateSoftSkills(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	customLabel := app.Profile.CustomSoftSkillLabel[lang]
	// TODO improve
	if app.JobPosition.SoftSkillsLabel != "" {
		customLabel = app.JobPosition.SoftSkillsLabel
	}
	defaultLabel := lang.SoftSkills()
	return convertListSkill(list, action, customLabel, defaultLabel, err, lang)
}

func lsaConvInterests(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	list, action, err := util.CalculateInterest(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	customLabel := app.Profile.CustomInterestLabel[lang]
	defaultLabel := lang.Interests()
	return convertListSkill(list, action, customLabel, defaultLabel, err, lang)
}

func lsaConvHobbies(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	list, action, err := util.CalculateHobbies(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	customLabel := app.Profile.CustomHobbiesLabel[lang]
	defaultLabel := lang.Hobbies()
	return convertListSkill(list, action, customLabel, defaultLabel, err, lang)
}

func convertListSkill(skills []genmodel.LeftSideAction, action genmodel.LeftSideActionType, customLabel string, defaultLabel string, err error, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction

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

func lsaConvTimeAmount(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	res.Label = app.Profile.CustomTimeAmountLabel[lang]
	if res.Label == "" {
		res.Label = lang.TimeAmount()
	}

	skills, action, err := util.CalculateTimeAmount(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	if err != nil {
		return res, action, false, err
	}

	//log.Printf("lsaConvTimeAmount 2 %d\n", len(skills))

	//maxSkills := maxProfessionalSkills
	//if maxProfessionalSkills+1 == len(skills) {
	//	maxSkills = maxProfessionalSkills + 1
	//}

	fulltime := 40.0
	endingText := "h"

	for _, skill := range skills {
		fulltime = 40.0
		endingText = "h"

		if skill.Full > 0.1 {
			fulltime = skill.Full
		}
		//log.Printf("lsaConvTimeAmount 3 %f - %f\n", skill.Full, fulltime)

		if skill.CurrencyEnding != "" {
			endingText = skill.CurrencyEnding
		}
		//log.Printf("lsaConvTimeAmount 4 %s - %s\n", skill.CurrencyEnding, endingText)

		delta := (skill.Max - skill.Min) / fulltime
		deltafull := 1 - (skill.Max / fulltime)

		mainlabel := lang.String(skill.SingleLabel)
		maxlabel := fmt.Sprintf("%.0f%s\n", skill.Max, endingText)
		if lang.String(skill.SingleLabel) == "" {
			mainlabel = fmt.Sprintf("%.0f%s\n", skill.Min, endingText)
		} else {
			maxlabel = ""
		}

		res.Range = append(res.Range, texmodel.RangeLsa{
			Name:            lang.String(skill.Name),
			Min:             skill.Min / fulltime,
			MinString:       fmt.Sprintf("%.3f", skill.Min/fulltime),
			MinLabel:        mainlabel,
			Max:             skill.Max / fulltime,
			MaxString:       fmt.Sprintf("%.3f", skill.Max/fulltime),
			MaxLabel:        maxlabel,
			DeltaMaxMin:     fmt.Sprintf("%f", delta),
			DeltaMaxMinHalf: fmt.Sprintf("%f", delta/2),
			DeltaMaxFull:    fmt.Sprintf("%.3f", deltafull),
		})
	}

	return res, action, len(skills) > 0, nil
}

func lsaConvMoneyAmount(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	res.Label = app.Profile.CustomMoneyAmountLabel[lang]
	if res.Label == "" {
		res.Label = lang.MoneyAmount()
	}

	skills, action, err := util.CalculateMoneyAmount(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)
	if err != nil {
		return res, action, false, err
	}

	//maxSkills := maxProfessionalSkills
	//if maxProfessionalSkills+1 == len(skills) {
	//	maxSkills = maxProfessionalSkills + 1
	//}

	fullmoney := 100.0
	currencyText := "T€"

	for _, skill := range skills {
		if skill.Full > 0.0 {
			fullmoney = skill.Full
		}
		if skill.CurrencyEnding != "" {
			currencyText = skill.CurrencyEnding
		}

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

	return res, action, len(skills) > 0, nil
}

func lsaConvPublication(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	action := genmodel.Lsa_publication
	if !lsaTypeIsActive(app, action) {
		return res, action, false, nil
	}

	res.Label = app.Profile.CustomPublicationLabel[lang]
	if res.Label == "" {
		res.Label = lang.Publication()
	}

	publications := filterPublication(app)
	for _, pub := range publications {
		res.PublicationLsa = append(res.PublicationLsa, texmodel.PublicationLsa{
			Name:                lang.String(pub.Title),
			Time:                pub.Date,
			SubTitle:            lang.String(pub.SubTitle),
			Image:               pub.CoverImage,
			Description:         lang.String(pub.Content),
			DescriptionShortLsa: lang.String(pub.ContentShortLsa),
			Publisher:           lang.String(pub.Publisher),
			ExternalLinks:       pub.ExternalLinks,
			DocumentLinks:       pub.DocumentLinks,
		})
	}
	return res, action, len(publications) > 0, nil
}

func lsaConvAward(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	var res texmodel.LeftSideAction
	action := genmodel.Lsa_award
	if !lsaTypeIsActive(app, action) {
		return res, action, false, nil
	}

	res.Label = app.Profile.CustomAwardLabel[lang]
	if res.Label == "" {
		res.Label = lang.Award()
	}

	awards := filterAward(app)
	for _, award := range awards {
		res.AwardLsa = append(res.AwardLsa, texmodel.AwardLsa{
			Name:                lang.String(award.Title),
			Time:                award.Date,
			Description:         lang.String(award.Content),
			DescriptionShortLsa: lang.String(award.ContentShortLsa),
			ExternalLinks:       award.ExternalLinks,
			DocumentLinks:       award.DocumentLinks,
		})
	}
	return res, action, len(awards) > 0, nil
}

func lsaConvLanguage(app *genmodel.Application, lang lang.Language) (texmodel.LeftSideAction, genmodel.LeftSideActionType, bool, error) {
	langs, action, err := util.CalculateLanguage(app, app.JobPosition.TwoSideStyle.Skills, app.JobPosition.TwoSideStyle.RemoveSkills)

	var res texmodel.LeftSideAction
	res.Label = app.Profile.CustomLanguageLabel[lang]
	if res.Label == "" {
		res.Label = lang.Languages()
	}
	if err != nil {
		return res, action, false, err
	}

	for _, l := range langs {
		// TODO improve with something like 'and not mentioned explicitly'
		//if l.DefaultShow == "no" {
		//	continue
		//}
		res.Languages = append(res.Languages, texmodel.Language{
			Name:           lang.String(l.Name),
			Level:          lang.String(l.Level),
			ShowRefExplain: l.ShowRefExplain,
			DocumentLinks:  l.DocumentLinks,
		})

		if l.ShowRefExplain {
			res.ShowLanguageRef = true
		}
	}
	return res, action, len(langs) > 0, nil
}

func lsaTypeIsActive(app *genmodel.Application, action genmodel.LeftSideActionType) bool {
	style := style(app)
	if _, ok := style.SideOneLSATypes.Index(action); ok {
		return true
	}
	if _, ok := style.SideTwoLSATypes.Index(action); ok {
		return true
	}
	if _, ok := app.Profile.LeftSideActionType.Index(action); ok {
		return true
	}
	return false
}
