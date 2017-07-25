package maprost

import (
	"github.com/maprost/application/example/maprost/internal"
	"github.com/maprost/application/generator/genmodel"
)

func Application(company string) genmodel.Application {
	application := genmodel.Application{}
	application.Profile = internal.Profile()

	switch company {
	case "next":
		application.JobPosition = internal.Google()
	}

	return application
}
