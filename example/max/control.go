package max

import (
	"github.com/maprost/application/example/max/company/bunny"
	"github.com/maprost/application/example/max/company/santa"
	"github.com/maprost/application/example/max/profile"
	"github.com/maprost/application/generator/genmodel"
)

func Profile(company string) genmodel.Profile {
	prf := profile.New()

	switch company {
	case "santa":
		prf.Company = santa.New()
	case "bunny":
		prf.Company = bunny.New()
	}

	return prf
}
