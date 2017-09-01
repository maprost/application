package util

import (
	"github.com/maprost/application/generator/genmodel"
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
		application.JobPosition.FutureExperience.FutureExperience = true
		application.Profile.Experience = append([]genmodel.Experience{application.JobPosition.FutureExperience}, application.Profile.Experience...)
	}
}
