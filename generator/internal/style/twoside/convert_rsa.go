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

type texmodelRightSideAction struct {
	texmodel.RSA
	action   genmodel.RightSideActionType
	orderIdx int
}

func rsaFirstSide(x genmodel.RightSideActionType) bool {
	switch x {
	case genmodel.Rsa_profile,
		genmodel.Rsa_myMotivation,
		genmodel.Rsa_mainQuestion,
		genmodel.Rsa_experience,
		genmodel.Rsa_publication:
		return true
	default:
		return false
	}
}

func addRSA(data *texmodel.Index, app *genmodel.Application, local lang.Language) {
	style := style(app)
	if len(app.Profile.RightSideActionType) == 0 {
		app.Profile.RightSideActionType = genmodel.RightSideActionTypes{
			genmodel.Rsa_profile,
			genmodel.Rsa_myMotivation,
			genmodel.Rsa_mainQuestion,
			genmodel.Rsa_experience,
			genmodel.Rsa_publication,
			genmodel.Rsa_education,
			genmodel.Rsa_award,
		}
	}

	var sideOneRsa []texmodelRightSideAction
	var sideTwoRsa []texmodelRightSideAction
	addRightSideAction := func(fn func(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error)) {
		rsa, action, ok, err := fn(app, local)
		if err != nil {
			fmt.Println("rsa err: " + err.Error())
			return
		}
		if !ok {
			fmt.Println("rsa not ok: " + action.String())
			return
		}
		if len(rsa.TexList) == 0 {
			fmt.Println("rsa empty: " + action.String())
			return
		}

		if len(style.SideOneRSATypes) > 0 || len(style.SideTwoRSATypes) > 0 {
			if idx, ok := style.SideOneRSATypes.Index(action); ok {
				sideOneRsa = append(sideOneRsa, texmodelRightSideAction{
					RSA:      rsa,
					action:   action,
					orderIdx: idx,
				})
			} else if idx, ok := style.SideTwoRSATypes.Index(action); ok {
				sideTwoRsa = append(sideTwoRsa, texmodelRightSideAction{
					RSA:      rsa,
					action:   action,
					orderIdx: idx,
				})
			} else {
				// do nothing
			}
		} else {
			if idx, ok := app.Profile.RightSideActionType.Index(action); ok {
				if rsaFirstSide(action) {
					sideOneRsa = append(sideOneRsa, texmodelRightSideAction{
						RSA:      rsa,
						action:   action,
						orderIdx: idx,
					})
				} else {
					sideTwoRsa = append(sideTwoRsa, texmodelRightSideAction{
						RSA:      rsa,
						action:   action,
						orderIdx: idx,
					})
				}
			}
		}
	}

	addRightSideAction(convertProfile)
	addRightSideAction(convertMyMotivation)
	addRightSideAction(convertMainQuestion)
	addRightSideAction(convertExperience)
	addRightSideAction(convertEducation)
	addRightSideAction(convertPublication)
	addRightSideAction(convertAward)

	conv := func(s []texmodelRightSideAction) []texmodel.RSA {
		sort.Slice(s, func(i, j int) bool {
			return s[i].orderIdx < s[j].orderIdx
		})

		var res []texmodel.RSA
		for _, s := range s {
			res = append(res, s.RSA)
		}
		return res
	}
	data.SideOneRSA = conv(sideOneRsa)
	data.SideTwoRSA = conv(sideTwoRsa)
	data.SideOneRSA, data.SideTwoRSA = splitRSA(app, data.SideOneRSA, data.SideTwoRSA)
}

func convertProfile(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	return convertTxt(
		customCustomDefaultString(app.JobPosition.ProfileLabel, app.Profile.CustomProfilTextLabel, local, local.Profile()),
		util.DefaultValue(app.JobPosition.ProfileText, local.String(app.Profile.GeneralProfileText)),
		genmodel.Rsa_profile,
	)
}

func convertMyMotivation(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	return convertTxt(
		customCustomDefaultString(app.JobPosition.MyMotivationLabel, app.Profile.CustomMyMotivationTextLabel, local, local.Motivation()),
		util.DefaultValue(app.JobPosition.MyMotivationText, local.String(app.Profile.GeneralMyMotivationText)),
		genmodel.Rsa_myMotivation,
	)
}

func convertMainQuestion(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	return convertTxt(
		customCustomDefaultString(app.JobPosition.MainQuestionText, app.Profile.CustomMainQuestionTextLabel, local, local.MainQuestion()),
		util.DefaultValue(app.JobPosition.MainQuestionText, local.String(app.Profile.GeneralMainQuestionText)),
		genmodel.Rsa_mainQuestion,
	)
}

func convertTxt(label string, txt string, action genmodel.RightSideActionType) (texmodel.RSA, genmodel.RightSideActionType, bool, error) {
	return texmodel.RSA{
		Label:   label,
		TexList: []string{txt + "\\bigskip"},
	}, action, txt != "", nil
}

func convertExperience(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	style := style(app)

	convTransLang := func(typ genmodel.ExperiencePart, tm ...lang.TranslationMap) string {
		for _, r := range style.RemoveExperiencePart {
			if r == typ {
				return ""
			}
		}
		var idx int
		if i, ok := style.ShowExperiencePart[typ]; ok {
			idx = i - 1
		} else if style.ShowExperienceParts > 0 {
			idx = style.ShowExperienceParts - 1
		} else {
			idx = 0
		}

		t := tm[idx]
		if len(t) == 0 {
			return ""
		}
		return local.String(t)
	}

	for i, exp := range app.Profile.Experience {
		if len(style.Experience) > 0 {
			if !findId(exp.Id, style.Experience) {
				continue
			}
		}
		if findId(exp.Id, style.RemoveExperience) {
			continue
		}

		timeRange := convertTime(exp.StartTime, exp.EndTime, local)
		if i == 0 && exp.FutureExperience {
			timeRange = local.PossibleAt() + "~~" + exp.StartTime
		}

		res := texmodel.Experience{
			Position:      convTransLang(genmodel.ExperiencePart_jobPosition, exp.JobPosition, exp.JobPosition2, exp.JobPosition3),
			Description:   convTransLang(genmodel.ExperiencePart_description, exp.Description, exp.Description2, exp.Description3),
			Project:       convTransLang(genmodel.ExperiencePart_project, exp.Project, exp.Project2, exp.Project3),
			Role:          convTransLang(genmodel.ExperiencePart_role, exp.Role, exp.Role2, exp.Role3),
			Company:       exp.Company,
			Tech:          convTransLang(genmodel.ExperiencePart_techStack, exp.TechStack, exp.TechStack2, exp.TechStack3),
			Time:          timeRange,
			QuitReason:    local.String(exp.QuitReason),
			DocumentLinks: exp.DocumentLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "experience.tex", res)
		if err != nil {
			return texmodel.RSA{}, 0, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasTechStack = rsa.HasTechStack || len(res.Tech) > 0
		rsa.HasProjects = rsa.HasProjects || len(res.Project) > 0
		rsa.HasRole = rsa.HasRole || len(res.Role) > 0
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
	}

	rsa.Label = customDefaultString(app.Profile.CustomExperienceLabel, local, local.Experience())
	return rsa, genmodel.Rsa_experience, true, nil
}

func convertEducation(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	style := style(app)

	for _, edu := range app.Profile.Education {
		if len(style.Education) > 0 {
			if !findId(edu.Id, style.Education) {
				continue
			}
		}
		if findId(edu.Id, style.RemoveEducation) {
			continue
		}

		res := texmodel.Education{
			Graduation:    local.String(edu.Graduation),
			FinalGrade:    local.String(edu.FinalGrade),
			QuitReason:    local.String(edu.QuitReason),
			Institute:     edu.Institute,
			Focus:         local.String(edu.Focus),
			Time:          convertTime(edu.StartTime, edu.EndTime, local),
			DocumentLinks: edu.DocumentLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "education.tex", res)
		if err != nil {
			return texmodel.RSA{}, 0, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
	}

	rsa.Label = customDefaultString(app.Profile.CustomEducationLabel, local, local.Education())
	return rsa, genmodel.Rsa_education, true, nil
}

func convertPublication(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	action = genmodel.Rsa_publication
	if lsaTypeIsActive(app, genmodel.Lsa_publication) {
		return texmodel.RSA{}, action, false, nil
	}

	for _, pub := range filterPublication(app) {
		res := texmodel.Publication{
			Title:         local.String(pub.Title),
			Publisher:     local.String(pub.Publisher),
			Time:          convertTime(pub.Date, pub.Date, local),
			Content:       "",
			DocumentLinks: pub.DocumentLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "publication.tex", res)
		if err != nil {
			return texmodel.RSA{}, 0, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
	}

	rsa.Label = customDefaultString(app.Profile.CustomPublicationLabel, local, local.Publication())
	return rsa, action, true, nil
}

func filterPublication(app *genmodel.Application) []genmodel.Publication {
	style := style(app)
	var res []genmodel.Publication
	for _, pub := range app.Profile.Publication {
		if len(style.Publication) > 0 {
			if !findId(pub.Id, style.Publication) {
				continue
			}
		}
		if findId(pub.Id, style.RemovePublication) {
			continue
		}
		res = append(res, pub)
	}
	return res
}

func convertAward(app *genmodel.Application, local lang.Language) (rsa texmodel.RSA, action genmodel.RightSideActionType, ok bool, err error) {
	action = genmodel.Rsa_award
	if lsaTypeIsActive(app, genmodel.Lsa_award) {
		return texmodel.RSA{}, action, false, nil
	}

	for _, awa := range filterAward(app) {
		res := texmodel.Award{
			Title:           local.String(awa.Title),
			Institute:       local.String(awa.Institute),
			Time:            awa.Date,
			Content:         local.String(awa.Content),
			ContentShortLsa: local.String(awa.ContentShortLsa),
			DocumentLinks:   awa.DocumentLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "award.tex", res)
		if err != nil {
			return texmodel.RSA{}, action, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
	}

	rsa.Label = customDefaultString(app.Profile.CustomAwardLabel, local, local.Award())
	return rsa, action, true, nil
}

func filterAward(app *genmodel.Application) []genmodel.Award {
	style := style(app)
	var res []genmodel.Award
	for _, pub := range app.Profile.Award {
		if len(style.Award) > 0 {
			if !findId(pub.Id, style.Award) {
				continue
			}
		}
		if findId(pub.Id, style.RemoveAward) {
			continue
		}
		res = append(res, pub)
	}
	return res
}

func convertTime(start string, end string, lang lang.Language) string {
	if end == start {
		return start
	}
	if end == "" {
		return lang.Since() + "~~" + start
	}
	return start + " - " + end
}

func splitRSA(app *genmodel.Application, sideOne []texmodel.RSA, sideTwo []texmodel.RSA) ([]texmodel.RSA, []texmodel.RSA) {
	style := style(app)
	if style.SideOneRSAItems > 0 {
		findIdx := func(l []texmodel.RSA, maxIdx int) (int, int, bool, int) {
			count := 0
			rsaIdx := 0
			rsaListIdx := 0
			found := false
			for i, rsa := range l {
				if found {
					break
				}
				fmt.Println("check rsa count: ", count, maxIdx)
				if count == maxIdx {
					fmt.Println("found rsa")
					rsaIdx = i
					rsaListIdx = 0
					found = true
					break
				}
				for j := range rsa.TexList {
					fmt.Println("check texList count: ", count, maxIdx)
					if count == maxIdx {
						fmt.Println("found texList")
						rsaIdx = i
						rsaListIdx = j
						found = true
						break
					}
					count++
				}
			}
			return rsaIdx, rsaListIdx, found, count
		}

		sideOneRSAIdx, sideOneRSAListIdx, found, count := findIdx(sideOne, style.SideOneRSAItems)
		if found {
			return splitSideOneRSA(sideOne, sideTwo, sideOneRSAIdx, sideOneRSAListIdx)
		}

		sideTwoRSAIdx, sideTwoRSAListIdx, found, _ := findIdx(sideTwo, style.SideOneRSAItems-count)
		if found {
			return splitSideTwoRSA(sideOne, sideTwo, sideTwoRSAIdx, sideTwoRSAListIdx)
		} else {
			// if not found
			return append(sideOne, sideTwo...), nil
		}
	}

	return sideOne, sideTwo
}

func splitSideOneRSA(sideOne []texmodel.RSA, sideTwo []texmodel.RSA, rsaIdx int, rsaListIdx int) ([]texmodel.RSA, []texmodel.RSA) {
	fmt.Println("splitSideOneRSA rsaIdx: ", rsaIdx, " -- rsaListIdx: ", rsaListIdx)

	var newSideOne []texmodel.RSA
	var newSideTwo []texmodel.RSA
	if rsaIdx > 0 {
		newSideOne = sideOne[:rsaIdx]
	}

	rsa := sideOne[rsaIdx]
	if rsaListIdx == 0 {
		newSideTwo = append(newSideTwo, rsa)
	} else {
		// split textList
		newSideOne = append(newSideOne, texmodel.RSA{
			Label:   rsa.Label,
			TexList: rsa.TexList[:rsaListIdx],
		})
		newSideTwo = append(newSideTwo, texmodel.RSA{
			// no label
			TexList: rsa.TexList[rsaListIdx:],
		})
	}

	if rsaIdx < len(sideOne) {
		newSideTwo = append(newSideTwo, sideOne[rsaIdx+1:]...)
	}
	newSideTwo = append(newSideTwo, sideTwo...)
	return newSideOne, newSideTwo
}

func splitSideTwoRSA(sideOne []texmodel.RSA, sideTwo []texmodel.RSA, rsaIdx int, rsaListIdx int) ([]texmodel.RSA, []texmodel.RSA) {
	fmt.Println("splitSideTwoRSA rsaIdx: ", rsaIdx, " -- rsaListIdx: ", rsaListIdx)

	var newSideOne []texmodel.RSA
	newSideOne = append(newSideOne, sideOne...)

	var newSideTwo []texmodel.RSA
	if rsaIdx > 0 {
		newSideOne = append(newSideOne, sideTwo[:rsaIdx]...)
	}

	rsa := sideTwo[rsaIdx]
	if rsaListIdx == 0 {
		newSideTwo = append(newSideTwo, rsa)
	} else {
		// split textList
		newSideOne = append(newSideOne, texmodel.RSA{
			Label:   rsa.Label,
			TexList: rsa.TexList[:rsaListIdx],
		})
		newSideTwo = append(newSideTwo, texmodel.RSA{
			// no label
			TexList: rsa.TexList[rsaListIdx:],
		})
	}

	if rsaIdx < len(sideTwo) {
		newSideTwo = append(newSideTwo, sideTwo[rsaIdx+1:]...)
	}

	return newSideOne, newSideTwo
}
