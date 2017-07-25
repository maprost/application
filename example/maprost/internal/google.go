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
		MainColor:      "800000",
	}
}
