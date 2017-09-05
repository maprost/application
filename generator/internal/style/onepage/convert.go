package onepage

import (
	"strings"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/onepage/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func initData(application *genmodel.Application, lang lang.Language) (data texmodel.Index, err error) {
	profSkills, otherProfSkills, err := convertProfSkills(application)
	if err != nil {
		return
	}

	softSkills, err := convertSoftSkills(application)
	if err != nil {
		return
	}

	aboutMe := util.DefaultValue(application.JobPosition.MotivationText, application.Profile.GeneralMotivationText)

	data = texmodel.Index{
		Label:           lang,
		MainColor:       util.DefaultColor(application.JobPosition.MainColor),
		Name:            util.JoinStrings(application.Profile.FirstName, " ", application.Profile.LastName),
		Title:           application.JobPosition.Title,
		Image:           util.DefaultImage(application.Profile.Image),
		Email:           application.Profile.Email,
		Phone:           application.Profile.Phone,
		Location:        util.JoinStrings(application.Profile.Address.City, ", ", application.Profile.Address.Country),
		Nationality:     application.Profile.Nationality,
		Websites:        convertWebsites(application),
		OtherProfSkills: otherProfSkills,
		ProfSkills:      profSkills,
		SoftSkills:      softSkills,
		Hobbies:         strings.Join(application.Profile.Hobbies, ", "),
		Interest:        strings.Join(application.Profile.Interest, ", "),
		Language:        convertLanguage(application),
		AboutMe:         aboutMe,
		Experience:      convertExperience(application, lang),
		Education:       convertEducation(application, lang),
	}
	return
}

func convertProfSkills(application *genmodel.Application) (profSkills []texmodel.Skill, otherProfSkills string, err error) {
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
				Name:   skill.Name,
				Rating: skill.Rating,
			})
		} else {
			if i > maxSkills {
				otherProfSkills += ", "
			}
			otherProfSkills += skill.Name
		}
	}
	return
}

func convertSoftSkills(application *genmodel.Application) (softSkills string, err error) {
	skills, err := util.CalculateSoftSkills(application)
	if err != nil {
		return
	}

	for i, skill := range skills {
		if i > 0 {
			softSkills += ", "
		}
		softSkills += skill.Name
	}
	return
}

func convertLanguage(application *genmodel.Application) (out []texmodel.Language) {
	for _, language := range application.Profile.Language {
		out = append(out, texmodel.Language{
			Name:  language.Name,
			Level: language.Level,
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
			Position:    exp.JobPosition,
			Description: util.ReplaceNewLine(exp.Description),
			Company:     exp.Company,
			Tech:        exp.TechStack,
			Time:        timeRange,
		})
	}
	return
}

func convertEducation(application *genmodel.Application, lang lang.Language) (education []texmodel.Education) {
	for _, edu := range application.Profile.Education {
		graduationWithGrade := edu.Graduation
		if edu.FinalGrade != "" {
			graduationWithGrade = graduationWithGrade + " (" + edu.FinalGrade + ")"
		}

		education = append(education, texmodel.Education{
			Graduation: graduationWithGrade,
			Institute:  edu.Institute,
			Focus:      edu.Focus,
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
