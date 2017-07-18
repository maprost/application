package max

import (
	"github.com/maprost/application/example/max/jobposition/bunny"
	"github.com/maprost/application/example/max/jobposition/santa"
	"github.com/maprost/application/example/max/profile"
	"github.com/maprost/application/generator/genmodel"
)

func Application(company string) genmodel.Application {
	application := genmodel.Application{}
	application.Profile = profile.New()

	switch company {
	case "santa":
		application.JobPosition = santa.New()
	case "bunny":
		application.JobPosition = bunny.New()
	}

	return application
}
