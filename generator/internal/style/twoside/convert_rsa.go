package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/lang"
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
		} else {
			idx = style.ShowExperienceParts - 1
		}
		return local.String(tm[idx])
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
			appendSideOne(expRes)
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

func convertTime(start string, end string, lang lang.Language) string {
	if end == "" {
		return lang.Since() + "~~" + start
	}
	return start + " - " + end
}
