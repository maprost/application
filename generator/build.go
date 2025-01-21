package generator

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
	"github.com/maprost/application/generator/internal/util"
)

func Build(application genmodel.Application, outputPath string) (err error) {
	outputPath, file := generateOutput(outputPath, &application)
	err = build(application, outputPath, file)

	return
}

func BuildAndClean(application genmodel.Application, outputPath string) (err error) {
	outputPath, file := generateOutput(outputPath, &application)

	err = build(application, outputPath, file)
	if err != nil {
		return
	}

	err = compiler.CleanUp(outputPath, file)
	return
}

func generalConvert(application *genmodel.Application) {
	util.JoinExperience(application)
}

func generateOutput(outputPath string, application *genmodel.Application) (path string, file string) {
	path = application.JobPosition.OutputPath
	if path == "" {
		path = outputPath
	}

	file = application.JobPosition.FileName
	if file == "" {
		file = "application"
	}

	return path, file
}

func build(application genmodel.Application, outputPath string, file string) (err error) {
	var data interface{}
	var path string
	var mainFile string
	var subFiles []string

	generalConvert(&application)

	data, err = application.JobPosition.Style.Data(&application)
	if err != nil {
		return
	}
	path, mainFile, subFiles = application.JobPosition.Style.Files()

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
