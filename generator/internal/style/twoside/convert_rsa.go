package twoside

import (
	"fmt"
	"log"
	"sort"
	"strings"

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
	label := ""
	//label += "\\colorbox{green}{"
	label += customCustomDefaultString(app.JobPosition.MainQuestionLabel, app.Profile.CustomMainQuestionTextLabel, local, local.MainQuestion())
	//label += "}"

	value := ""
	//value += "\\colorbox{green}{"
	value += util.DefaultValue(app.JobPosition.MainQuestionText, local.String(app.Profile.GeneralMainQuestionText))
	//value += "}"

	return convertTxt(
		label,
		value,
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
		// TODO improve with something like 'and not mentioned explicitly'
		//if exp.DefaultShow == "no" {
		//	continue
		//}
		if !exp.FutureExperience {
			if len(style.Experience) > 0 {
				if !findId(exp.Id, style.Experience) {
					continue
				}
			}
			if findId(exp.Id, style.RemoveExperience) {
				continue
			}
			if exp.MustBeSelected && len(style.Experience) == 0 {
				continue
			}
		}

		timeRange := convertTime(exp.StartTime, exp.EndTime, local)
		timeRange2 := ""
		if i == 0 && exp.FutureExperience {
			timeRange = local.PossibleAt() + "~~" + exp.StartTime
			if exp.EndTime != "" {
				timeRange2 += local.PossibleUntil() + " " + exp.EndTime
			}
		}

		// TODO improve
		if local.String(exp.JobPosition2) == "" && local.String(exp.JobPosition) != "" {
			exp.JobPosition2 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.JobPosition2)) == "xxx" {
			exp.JobPosition2 = lang.DefaultTranslation("")
		}

		if local.String(exp.JobPosition3) == "" && local.String(exp.JobPosition) != "" {
			exp.JobPosition3 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.JobPosition3)) == "xxx" {
			exp.JobPosition3 = lang.DefaultTranslation("")
		}

		if local.String(exp.Description2) == "" && local.String(exp.Description) != "" {
			exp.Description2 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.Description2)) == "xxx" {
			exp.Description2 = lang.DefaultTranslation("")
		}

		if local.String(exp.Description3) == "" && local.String(exp.Description) != "" {
			exp.Description3 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.Description3)) == "xxx" {
			exp.Description3 = lang.DefaultTranslation("")
		}

		if local.String(exp.Role2) == "" && local.String(exp.Role) != "" {
			exp.Role2 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.Role2)) == "xxx" {
			exp.Role2 = lang.DefaultTranslation("")
		}

		if local.String(exp.Role3) == "" && local.String(exp.Role) != "" {
			exp.Role3 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.Role3)) == "xxx" {
			exp.Role3 = lang.DefaultTranslation("")
		}

		if local.String(exp.Project2) == "" && local.String(exp.Project) != "" {
			exp.Project2 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.Project2)) == "xxx" {
			exp.Project2 = lang.DefaultTranslation("")
		}

		if local.String(exp.Project3) == "" && local.String(exp.Project) != "" {
			exp.Project3 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.Project3)) == "xxx" {
			exp.Project3 = lang.DefaultTranslation("")
		}

		if local.String(exp.TechStack2) == "" && local.String(exp.TechStack) != "" {
			exp.TechStack2 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.TechStack2)) == "xxx" {
			exp.TechStack2 = lang.DefaultTranslation("")
		}

		if local.String(exp.TechStack3) == "" && local.String(exp.TechStack) != "" {
			exp.TechStack3 = lang.DefaultTranslation("wie 1")
		} else if strings.ToLower(local.String(exp.TechStack3)) == "xxx" {
			exp.TechStack3 = lang.DefaultTranslation("")
		}

		if local.String(exp.JobPosition) != "" && strings.ToLower(local.String(exp.JobPosition2)) == "wie 1" {
			log.Printf("wie 11111 %+v", exp.JobPosition)
			exp.JobPosition2 = exp.JobPosition
			log.Printf("wie 11111_2 %+v", exp.JobPosition2)
		}
		if local.String(exp.JobPosition) != "" && strings.ToLower(local.String(exp.JobPosition3)) == "wie 1" {
			exp.JobPosition3 = exp.JobPosition
		}

		if local.String(exp.Description) != "" && strings.ToLower(local.String(exp.Description2)) == "wie 1" {
			exp.Description2 = exp.Description
		}
		if local.String(exp.Description) != "" && strings.ToLower(local.String(exp.Description3)) == "wie 1" {
			exp.Description3 = exp.Description
		}

		if local.String(exp.Project) != "" && strings.ToLower(local.String(exp.Project2)) == "wie 1" {
			log.Printf("wie 11111 %+v", exp.Project)
			exp.Project2 = exp.Project
			log.Printf("wie 11111_2 %+v", exp.Project2)
		}
		if local.String(exp.Project) != "" && strings.ToLower(local.String(exp.Project3)) == "wie 1" {
			exp.Project3 = exp.Project
		}

		if local.String(exp.Role) != "" && strings.ToLower(local.String(exp.Role2)) == "wie 1" {
			exp.Role2 = exp.Role
		}
		if local.String(exp.Role) != "" && strings.ToLower(local.String(exp.Role3)) == "wie 1" {
			exp.Role3 = exp.Role
		}

		if local.String(exp.TechStack) != "" && strings.ToLower(local.String(exp.TechStack2)) == "wie 1" {
			exp.TechStack2 = exp.TechStack
		}
		if local.String(exp.TechStack) != "" && strings.ToLower(local.String(exp.TechStack3)) == "wie 1" {
			exp.TechStack3 = exp.TechStack
		}

		res := texmodel.Experience{
			Position:          convTransLang(genmodel.ExperiencePart_jobPosition, exp.JobPosition, exp.JobPosition2, exp.JobPosition3),
			Description:       convTransLang(genmodel.ExperiencePart_description, exp.Description, exp.Description2, exp.Description3),
			Project:           convTransLang(genmodel.ExperiencePart_project, exp.Project, exp.Project2, exp.Project3),
			Role:              convTransLang(genmodel.ExperiencePart_role, exp.Role, exp.Role2, exp.Role3),
			Company:           exp.Company,
			Tech:              convTransLang(genmodel.ExperiencePart_techStack, exp.TechStack, exp.TechStack2, exp.TechStack3),
			Time:              timeRange,
			TimeSecondLine:    timeRange2,
			QuitReason:        local.String(exp.QuitReason),
			DocumentLinks:     exp.DocumentLinks,
			ExternalLinks:     exp.ExternalLinks,
			FutureExperience:  exp.FutureExperience,
			FutureColorChange: exp.FutureColorChange,
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
		rsa.HasExternalLinks = rsa.HasExternalLinks || len(res.ExternalLinks) > 0
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
		if edu.MustBeSelected && len(style.Education) == 0 {
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
			ExternalLinks: edu.ExternalLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "education.tex", res)
		if err != nil {
			return texmodel.RSA{}, 0, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
		rsa.HasExternalLinks = rsa.HasExternalLinks || len(res.ExternalLinks) > 0
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
			ExternalLinks: pub.ExternalLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "publication.tex", res)
		if err != nil {
			return texmodel.RSA{}, 0, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
		rsa.HasExternalLinks = rsa.HasExternalLinks || len(res.ExternalLinks) > 0
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
		if pub.MustBeSelected && len(style.Publication) == 0 {
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
			ExternalLinks:   awa.ExternalLinks,
		}

		resTex, err := compiler.CompileSubTex(templatePath(), "award.tex", res)
		if err != nil {
			return texmodel.RSA{}, action, false, err
		}

		rsa.TexList = append(rsa.TexList, resTex)
		rsa.HasDocumentLinks = rsa.HasDocumentLinks || len(res.DocumentLinks) > 0
		rsa.HasExternalLinks = rsa.HasExternalLinks || len(res.ExternalLinks) > 0
	}

	rsa.Label = customDefaultString(app.Profile.CustomAwardLabel, local, local.Award())
	return rsa, action, true, nil
}

func filterAward(app *genmodel.Application) []genmodel.Award {
	style := style(app)
	var res []genmodel.Award
	for _, award := range app.Profile.Award {
		if len(style.Award) > 0 {
			if !findId(award.Id, style.Award) {
				continue
			}
		}
		if findId(award.Id, style.RemoveAward) {
			continue
		}
		if award.MustBeSelected && len(style.Award) == 0 {
			continue
		}
		res = append(res, award)
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
		sOne, sTwo := rsa.Split(rsaListIdx)
		newSideOne = append(newSideOne, sOne)
		newSideTwo = append(newSideTwo, sTwo)
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
		sOne, sTwo := rsa.Split(rsaListIdx)
		newSideOne = append(newSideOne, sOne)
		newSideTwo = append(newSideTwo, sTwo)
	}

	if rsaIdx < len(sideTwo) {
		newSideTwo = append(newSideTwo, sideTwo[rsaIdx+1:]...)
	}

	return newSideOne, newSideTwo
}
