package july

import (
	"github.com/maprost/application/generator/genmodel"
)

func Profile(company string) genmodel.Profile {
	return genmodel.Profile{
		Name:  "July Musertmann",
		Title: "Santa Clause",
	}
}
