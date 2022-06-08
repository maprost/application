package google

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
	"github.com/maprost/application/generator/style"
)

func Google() genmodel.JobPosition {
	return genmodel.JobPosition{
		MainColor: "4285F4",
		Style:     style.OneSide,
		Lang:      lang.German,
		Company:   "Google",
		Address: genmodel.JobAddress{
			Street:  "...",
			Zip:     "...",
			City:    "...",
			Country: "...",
		},
		Title:          "Software Defender",
		MotivationText: ``,
		FutureExperience: genmodel.FutureExperience{
			JobPosition: "Go Backend Developer",
			TechStack:   "Go, Docker",
			StartTime:   "Jan. 2018",
			Description: "",
		},
	}
}
