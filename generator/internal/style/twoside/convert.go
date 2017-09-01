package twoside

import (
	"strings"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/util"
)

func initIndex(application *genmodel.Application, shortVersion bool) (texmodel.Index, error) {
	return texmodel.Index{
		MainColor:    util.DefaultColor(application.JobPosition.MainColor),
		ShortVersion: shortVersion,
	}, nil
}

func createFirstPageData(application *genmodel.Application) (texmodel.FirstPage, error) {
	return texmodel.FirstPage{
		Name:  util.JoinStrings(application.Profile.FirstName, " ", application.Profile.LastName),
		Title: application.JobPosition.Title,
		Image: util.DefaultImage(application.Profile.Image),
	}, nil
}

func createCoverLetterData(application *genmodel.Application) (texmodel.CoverLetter, error) {
	return texmodel.CoverLetter{
		Name:          util.JoinStrings(application.Profile.FirstName, " ", application.Profile.LastName),
		Street:        application.Profile.Address.Street,
		Zip:           util.JoinStrings(application.Profile.Address.Zip, " ", application.Profile.Address.City),
		CompanyName:   application.JobPosition.Company,
		CompanyStreet: application.JobPosition.Address.Street,
		CompanyZip:    util.JoinStrings(application.JobPosition.Address.Zip, " ", application.JobPosition.Address.City),
		Text:          util.DefaultValue(application.JobPosition.MotivationText, application.Profile.GeneralMotivationText),
	}, nil
}

func createCVData(application *genmodel.Application, shortVersion bool) (data texmodel.CV, err error) {
	profSkills, otherProfSkills, err := convertProfSkills(application)
	if err != nil {
		return
	}

	softSkills, err := convertSoftSkills(application)
	if err != nil {
		return
	}

	aboutMe := ""
	if shortVersion {
		aboutMe = util.DefaultValue(application.JobPosition.MotivationText, application.Profile.GeneralMotivationText)
	}

	data = texmodel.CV{
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
		Experience:      convertExperience(application),
		Education:       convertEducation(application),
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
	for _, lang := range application.Profile.Language {
		out = append(out, texmodel.Language{
			Name:  lang.Name,
			Level: lang.Level,
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

func convertExperience(application *genmodel.Application) (experience []texmodel.Experience) {
	for i, exp := range application.Profile.Experience {

		timeRange := convertTime(exp.StartTime, exp.EndTime)
		if i == 0 && exp.FutureExperience {
			timeRange = "possible at~~" + exp.StartTime
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

func convertEducation(application *genmodel.Application) (education []texmodel.Education) {
	for _, edu := range application.Profile.Education {
		graduationWithGrade := edu.Graduation
		if edu.FinalGrade != "" {
			graduationWithGrade = graduationWithGrade + " (" + edu.FinalGrade + ")"
		}

		education = append(education, texmodel.Education{
			Graduation: graduationWithGrade,
			Institute:  edu.Institute,
			Focus:      edu.Focus,
			Time:       convertTime(edu.StartTime, edu.EndTime),
		})
	}
	return
}

func convertTime(start string, end string) string {
	if end == "" {
		return "since~~" + start
	}
	return start + " - " + end
}
