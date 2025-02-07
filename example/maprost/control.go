package maprost

import (
	"log"

	"github.com/maprost/application/example/maprost/internal"
	"github.com/maprost/application/example/maprost/internal/google"
	"github.com/maprost/application/generator/genmodel"
)

var allCompanies = map[string]func() genmodel.JobPosition{
	"google": google.Google,
}

func Application(company string) genmodel.Application {
	f, ok := allCompanies[company]
	if !ok {
		log.Fatalln("Can't find job position", company)
	}

	application := genmodel.Application{
		Profile:     internal.Profile(),
		JobPosition: f(),
	}

	return application
}
