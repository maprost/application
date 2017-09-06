package onepage

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/onepage/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func initData(application *genmodel.Application, language lang.Language) (data texmodel.Index, err error) {
	profSkills, otherProfSkills, err := convertProfSkills(application, language)
	if err != nil {
		return
	}

	softSkills, err := convertSoftSkills(application, language)
	if err != nil {
		return
	}

	aboutMe := util.DefaultValue(application.JobPosition.MotivationText, application.Profile.GeneralMotivationText)

	data = texmodel.Index{
		Label:           language,
		MainColor:       util.DefaultColor(application.JobPosition.MainColor),
		Name:            util.JoinStrings(application.Profile.FirstName, " ", application.Profile.LastName),
		Title:           application.JobPosition.Title,
		Image:           util.DefaultImage(application.Profile.Image),
		Email:           application.Profile.Email,
		Phone:           application.Profile.Phone,
		Location:        util.JoinStrings(application.Profile.Address.City, ", ", application.Profile.Address.Country),
		Nationality:     application.Profile.Nationality.String(language),
		Websites:        convertWebsites(application),
		OtherProfSkills: otherProfSkills,
		ProfSkills:      profSkills,
		SoftSkills:      softSkills,
		Hobbies:         lang.JoinTranslationMap(application.Profile.Hobbies, language, ", "),
		Interest:        lang.JoinTranslationMap(application.Profile.Interest, language, ", "),
		Language:        convertLanguage(application, language),
		AboutMe:         aboutMe,
		Experience:      convertExperience(application, language),
		Education:       convertEducation(application, language),
	}
	return
}

func convertProfSkills(application *genmodel.Application, lang lang.Language) (profSkills []texmodel.Skill, otherProfSkills string, err error) {
	skills, err := util.CalculateProfessionalSkills(application)
	if err != nil {
		return
	}

	maxSkills := maxProfessionalSkills
	if maxProfessionalSkills+1 == len(skills) {
		maxSkills = maxProfessionalSkills + 1
	}

	for i, skill := range skills {
		if i < maxSkills {
			profSkills = append(profSkills, texmodel.Skill{
				Name:   skill.Name.String(lang),
				Rating: skill.Rating,
			})
		} else {
			if i > maxSkills {
				otherProfSkills += ", "
			}
			otherProfSkills += skill.Name.String(lang)
		}
	}
	return
}

func convertSoftSkills(application *genmodel.Application, lang lang.Language) (softSkills string, err error) {
	skills, err := util.CalculateSoftSkills(application)
	if err != nil {
		return
	}

	for i, skill := range skills {
		if i > 0 {
			softSkills += ", "
		}
		softSkills += skill.Name.String(lang)
	}
	return
}

func convertLanguage(application *genmodel.Application, lang lang.Language) (out []texmodel.Language) {
	for _, language := range application.Profile.Language {
		out = append(out, texmodel.Language{
			Name:  language.Name.String(lang),
			Level: language.Level.String(lang),
		})
	}
	return
}

func convertWebsites(application *genmodel.Application) (websites []texmodel.Website) {
	for _, website := range application.Profile.Websites {
		websites = append(websites, texmodel.Website{
			Icon: util.WebsiteIcon(website),
			Url:  website,
		})
	}
	return
}

func convertExperience(application *genmodel.Application, lang lang.Language) (experience []texmodel.Experience) {
	for i, exp := range application.Profile.Experience {

		timeRange := convertTime(exp.StartTime, exp.EndTime, lang)
		if i == 0 && exp.FutureExperience {
			timeRange = lang.PossibleAt() + "~~" + exp.StartTime
		}

		experience = append(experience, texmodel.Experience{
			Position:    exp.JobPosition.String(lang),
			Description: util.ReplaceNewLine(exp.Description.String(lang)),
			Company:     exp.Company,
			Tech:        exp.TechStack.String(lang),
			Time:        timeRange,
		})
	}
	return
}

func convertEducation(application *genmodel.Application, lang lang.Language) (education []texmodel.Education) {
	for _, edu := range application.Profile.Education {
		graduationWithGrade := edu.Graduation.String(lang)
		if edu.FinalGrade != "" {
			graduationWithGrade = graduationWithGrade + " (" + edu.FinalGrade + ")"
		}

		education = append(education, texmodel.Education{
			Graduation: graduationWithGrade,
			Institute:  edu.Institute,
			Focus:      edu.Focus.String(lang),
			Time:       convertTime(edu.StartTime, edu.EndTime, lang),
		})
	}
	return
}

func convertTime(start string, end string, lang lang.Language) string {
	if end == "" {
		return lang.Since() + "~~" + start
	}
	return start + " - " + end
}
