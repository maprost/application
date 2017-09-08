package generator

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/application/generator/lang"
	"strings"
	"unicode"
)

func Build(application genmodel.Application, style Style, lang lang.Language, outputPath string) (err error) {
	var data interface{}
	var path string
	var mainFile string
	var subFiles []string

	generalConvert(&application)

	data, err = style.Data(&application, lang)
	if err != nil {
		return
	}
	path, mainFile, subFiles = style.Files()

	outputPath, file := generateOutput(outputPath, &application)

	err = compiler.CreateTexFile(outputPath, file, data, path, mainFile, subFiles...)
	if err != nil {
		return
	}

	err = compiler.Compile(outputPath, file)
	if err != nil {
		return
	}

	//err = compiler.CleanUp(outputPath, file)
	return
}

func generalConvert(application *genmodel.Application) {
	util.JoinExperience(application)
}

func generateOutput(outputPath string, application *genmodel.Application) (path string, file string) {
	// extract first word of company

	company := strings.TrimSpace(application.JobPosition.Company)
	index := 0
	for i, r := range application.JobPosition.Company {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			index = i
		} else {
			break
		}
	}

	return outputPath + "/" + strings.ToLower(company[:index+1]), "application"
}
