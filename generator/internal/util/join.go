package util

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
)

func JoinStrings(valueA string, sep string, valueB string) string {
	if valueA == "" {
		return valueB
	}
	if valueB == "" {
		return valueA
	}
	return valueA + sep + valueB
}

func JoinExperience(application *genmodel.Application) {
	empty := application.JobPosition.FutureExperience.Empty()
	if !empty {
		futureExp := genmodel.Experience{
			JobPosition:      lang.DefaultTranslation(application.JobPosition.FutureExperience.JobPosition),
			Description:      lang.DefaultTranslation(application.JobPosition.FutureExperience.Description),
			TechStack:        lang.DefaultTranslation(application.JobPosition.FutureExperience.TechStack),
			StartTime:        application.JobPosition.FutureExperience.StartTime,
			Company:          application.JobPosition.Company,
			FutureExperience: true,
		}

		application.Profile.Experience = append([]genmodel.Experience{futureExp}, application.Profile.Experience...)
	}
}
