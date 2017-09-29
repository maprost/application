package onepage

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/onepage/texmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
)

func initData(application *genmodel.Application) (data texmodel.Index, err error) {
	lang := application.JobPosition.Lang
	profSkills, otherProfSkills, err := convertProfSkills(application, lang)
	if err != nil {
		return
	}

	softSkills, err := convertSoftSkills(application, lang)
	if err != nil {
		return
	}

	aboutMe := util.DefaultValue(application.JobPosition.MotivationText, lang.String(application.Profile.GeneralMotivationText))

	data = texmodel.Index{
		Icon: texmodel.Icon{
			Project:   util.ProjectIconPath,
			Role:      util.RoleIconPath,
			TechStack: util.TechStackIconPath,
			Document:  util.DocumentIconPath,
		},
		Label:           lang,
		MainColor:       util.DefaultColor(application.JobPosition.MainColor),
		Name:            util.JoinStrings(application.Profile.FirstName, " ", application.Profile.LastName),
		Title:           util.DefaultValue(application.JobPosition.Title, application.Profile.Title),
		Image:           util.DefaultImage(application.Profile.Image),
		Email:           application.Profile.Email,
		Phone:           application.Profile.Phone,
		Location:        util.JoinStrings(lang.String(application.Profile.Address.City), ", ", lang.String(application.Profile.Address.Country)),
		Nationality:     lang.String(application.Profile.Nationality),
		Websites:        convertWebsites(application),
		OtherProfSkills: otherProfSkills,
		ProfSkills:      profSkills,
		SoftSkills:      softSkills,
		Hobbies:         lang.Join(application.Profile.Hobbies, ", "),
		Interest:        lang.Join(application.Profile.Interest, ", "),
		Language:        convertLanguage(application, lang),
		AboutMe:         aboutMe,
		Experience:      convertExperience(application, lang),
		Education:       convertEducation(application, lang),
		Attachment:      util.DefaultStringArray(application.JobPosition.Attachment, application.Profile.Attachment),
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
				Name:   lang.String(skill.Name),
				Rating: skill.Rating,
			})
		} else {
			if i > maxSkills {
				otherProfSkills += ", "
			}
			otherProfSkills += lang.String(skill.Name)
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
		softSkills += lang.String(skill.Name)
	}
	return
}

func convertLanguage(application *genmodel.Application, lang lang.Language) (out []texmodel.Language) {
	for _, language := range application.Profile.Language {
		out = append(out, texmodel.Language{
			Name:  lang.String(language.Name),
			Level: lang.String(language.Level),
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
			Position:      lang.String(exp.JobPosition),
			Description:   lang.String(exp.Description),
			Project:       lang.String(exp.Project),
			Role:          lang.String(exp.Role),
			Company:       exp.Company,
			Tech:          lang.String(exp.TechStack),
			Time:          timeRange,
			DocumentLinks: exp.DocumentLinks,
		})
	}
	return
}

func convertEducation(application *genmodel.Application, lang lang.Language) (education []texmodel.Education) {
	for _, edu := range application.Profile.Education {
		education = append(education, texmodel.Education{
			Graduation:    lang.String(edu.Graduation),
			FinalGrade:    lang.String(edu.FinalGrade),
			Institute:     edu.Institute,
			Focus:         lang.String(edu.Focus),
			Time:          convertTime(edu.StartTime, edu.EndTime, lang),
			DocumentLinks: edu.DocumentLinks,
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
