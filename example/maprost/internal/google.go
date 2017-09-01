package internal

import (
	"github.com/maprost/application/generator/genmodel"
)

func Google() genmodel.JobPosition {
	return genmodel.JobPosition{
		Company: "Google",
		Address: genmodel.Address{
			Street:  "...",
			Zip:     "...",
			City:    "...",
			Country: "...",
		},
		Title:          "Software Defender",
		MotivationText: ``,
		MainColor:      "4285F4",
		FutureExperience: genmodel.Experience{
			JobPosition: "Go Backend Developer",
			TechStack:   "Go, Docker",
			Company:     "Google",
			StartTime:   "Jan. 2018",
			EndTime:     "",
			Description: `

			`,
		},
	}
}
