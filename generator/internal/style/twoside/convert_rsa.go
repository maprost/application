package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/lang"
)

func addExperiences(data *texmodel.Index, app *genmodel.Application, lang lang.Language) {
	style := app.JobPosition.TwoSideStyle

	var sideOneExp []texmodel.Experience
	var sideTwoExp []texmodel.Experience
	for i, exp := range app.Profile.Experience {
		if len(style.Experience) > 0 {
			found := false
			for _, eId := range style.Experience {
				if exp.Id == eId {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		timeRange := convertTime(exp.StartTime, exp.EndTime, lang)
		if i == 0 && exp.FutureExperience {
			timeRange = lang.PossibleAt() + "~~" + exp.StartTime
		}

		expRes := texmodel.Experience{
			Position:      lang.String(exp.JobPosition),
			Description:   lang.String(exp.Description),
			Project:       lang.String(exp.Project),
			Role:          lang.String(exp.Role),
			Company:       exp.Company,
			Tech:          lang.String(exp.TechStack),
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

	label := customDefaultString(app.Profile.CustomExperienceLabel, lang, lang.Experience())
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
			found := false
			for _, eId := range style.Education {
				if edu.Id == eId {
					found = true
					break
				}
			}
			if !found {
				continue
			}
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
