package generator

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
	"github.com/maprost/application/generator/internal/util"
)

func Build(app genmodel.Application, outputPath string) (err error) {
	outputPath, file := generateOutput(outputPath, &app)
	err = build(app, outputPath, file)

	return
}

func BuildAndClean(app genmodel.Application, outputPath string) (err error) {
	outputPath, file := generateOutput(outputPath, &app)

	err = build(app, outputPath, file)
	if err != nil {
		return
	}

	err = compiler.CleanUp(outputPath, file)
	return
}

func generalConvert(app *genmodel.Application) {
	util.JoinExperience(app)
}

func generateOutput(outputPath string, app *genmodel.Application) (path string, file string) {
	return compiler.GenerateOutput(outputPath, app)
}

func build(app genmodel.Application, outputPath string, file string) (err error) {
	var data interface{}
	var path string
	var mainFile string
	var subFiles []string

	generalConvert(&app)

	data, err = app.JobPosition.Style.Data(&app, outputPath)
	if err != nil {
		return
	}
	path, mainFile, subFiles = app.JobPosition.Style.Files()

	err = compiler.CreateTexFile(outputPath, file, data, path, mainFile, subFiles...)
	if err != nil {
		return
	}

	err = compiler.Compile(outputPath, file)
	if err != nil {
		return
	}

	return
}
