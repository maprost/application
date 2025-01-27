package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/lang"
	"log"
)

func addExperiences(data *texmodel.Index, app *genmodel.Application, local lang.Language) {
	style := app.JobPosition.TwoSideStyle

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

	var sideOneExp []texmodel.Experience
	var sideTwoExp []texmodel.Experience
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

		expRes := texmodel.Experience{
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

		if style.SideOneExperienceSize > 0 {
			if len(sideOneExp) < style.SideOneExperienceSize {
				sideOneExp = append(sideOneExp, expRes)
			} else {
				sideTwoExp = append(sideTwoExp, expRes)
			}
		} else {
			sideOneExp = append(sideOneExp, expRes)
		}
	}

	data.SideOneExperience = sideOneExp
	data.SideTwoExperience = sideTwoExp

	label := customDefaultString(app.Profile.CustomExperienceLabel, local, local.Experience())
	if len(sideOneExp) > 0 {
		data.SideOneExperienceLabel = label
	} else if len(sideTwoExp) > 0 {
		data.SideTwoExperienceLabel = label
	}
}

func addEducation(data *texmodel.Index, app *genmodel.Application, lang lang.Language) {
	style := app.JobPosition.TwoSideStyle

	var sideOneEdu []texmodel.Education
	var sideTwoEdu []texmodel.Education

	appendSideOne := func(edu texmodel.Education) {
		if len(data.SideTwoExperience) > 0 {
			sideTwoEdu = append(sideTwoEdu, edu)
		} else {
			sideOneEdu = append(sideOneEdu, edu)
		}
	}

	for _, edu := range app.Profile.Education {
		if len(style.Education) > 0 {
			if !findId(edu.Id, style.Education) {
				continue
			}
		}
		if findId(edu.Id, style.RemoveEducation) {
			continue
		}

		expRes := texmodel.Education{
			Graduation:    lang.String(edu.Graduation),
			FinalGrade:    lang.String(edu.FinalGrade),
			QuitReason:    lang.String(edu.QuitReason),
			Institute:     edu.Institute,
			Focus:         lang.String(edu.Focus),
			Time:          convertTime(edu.StartTime, edu.EndTime, lang),
			DocumentLinks: edu.DocumentLinks,
		}

		if style.SideOneEducationSize > 0 {
			if len(sideOneEdu) < style.SideOneEducationSize {
				appendSideOne(expRes)
			} else {
				sideTwoEdu = append(sideTwoEdu, expRes)
			}
		} else {
			sideTwoEdu = append(sideTwoEdu, expRes)
		}
	}

	data.SideOneEducation = sideOneEdu
	data.SideTwoEducation = sideTwoEdu

	label := customDefaultString(app.Profile.CustomEducationLabel, lang, lang.Education())
	if len(sideOneEdu) > 0 {
		data.SideOneEducationLabel = label
	} else if len(sideTwoEdu) > 0 {
		data.SideTwoEducationLabel = label
	}
}

func addPublication(data *texmodel.Index, app *genmodel.Application, lang lang.Language) {
	style := app.JobPosition.TwoSideStyle

	log.Printf("addPublication 1")

	var sideOnePub []texmodel.Publication
	var sideTwoPub []texmodel.Publication

	appendSideOne := func(pub texmodel.Publication) {
		// TODO improve?
		if len(data.SideTwoExperience) > 0 {
			sideTwoPub = append(sideTwoPub, pub)
		} else {
			sideOnePub = append(sideOnePub, pub)
		}
	}

	for _, pub := range app.Profile.Publication {
		if len(style.Publication) > 0 {
			if !findId(pub.Id, style.Publication) {
				continue
			}
		}
		if findId(pub.Id, style.RemovePublication) {
			continue
		}

		expRes := texmodel.Publication{
			Title:         lang.String(pub.Title),
			Publisher:     lang.String(pub.Publisher),
			Time:          convertTime(pub.StartTime, pub.EndTime, lang),
			Content:       "",
			DocumentLinks: pub.DocumentLinks,
		}

		if style.SideOnePublicationSize > 0 {
			if len(sideOnePub) < style.SideOnePublicationSize {
				appendSideOne(expRes)
			} else {
				sideTwoPub = append(sideTwoPub, expRes)
			}
		} else {
			//sideTwoPub = append(sideTwoPub, expRes)
			appendSideOne(expRes)
		}
	}

	data.SideOnePublication = sideOnePub
	data.SideTwoPublication = sideTwoPub

	label := customDefaultString(app.Profile.CustomPublicationLabel, lang, lang.Publication())
	if len(sideOnePub) > 0 {
		data.SideOnePublicationLabel = label
	} else if len(sideTwoPub) > 0 {
		data.SideTwoPublicationLabel = label
	}

	log.Printf("addPublication 2")
	log.Printf("%+v", data)

}

func addAward(data *texmodel.Index, app *genmodel.Application, lang lang.Language) {
	style := app.JobPosition.TwoSideStyle

	var sideOnePub []texmodel.Award
	var sideTwoPub []texmodel.Award

	appendSideOne := func(awa texmodel.Award) {
		if len(data.SideTwoExperience) > 0 {
			sideTwoPub = append(sideTwoPub, awa)
		} else {
			sideOnePub = append(sideOnePub, awa)
		}
	}

	for _, awa := range app.Profile.Award {
		if len(style.Education) > 0 {
			if !findId(awa.Id, style.Award) {
				continue
			}
		}
		if findId(awa.Id, style.RemoveAward) {
			continue
		}

		expRes := texmodel.Award{
			Title:         lang.String(awa.Title),
			Institute:     lang.String(awa.Institute),
			Time:          convertTime(awa.StartTime, awa.EndTime, lang),
			Content:       "",
			DocumentLinks: awa.DocumentLinks,
		}

		if style.SideOneAwardSize > 0 {
			if len(sideOnePub) < style.SideOneAwardSize {
				appendSideOne(expRes)
			} else {
				sideTwoPub = append(sideTwoPub, expRes)
			}
		} else {
			sideTwoPub = append(sideTwoPub, expRes)
		}
	}

	data.SideOneAward = sideOnePub
	data.SideTwoAward = sideTwoPub

	label := customDefaultString(app.Profile.CustomAwardLabel, lang, lang.Award())
	if len(sideOnePub) > 0 {
		data.SideOneAwardLabel = label
	} else if len(sideTwoPub) > 0 {
		data.SideTwoAwardLabel = label
	}
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
