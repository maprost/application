package generator

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
	"github.com/maprost/application/generator/internal/util"
)

func Build(application genmodel.Application, style Style) (err error) {
	var data interface{}
	var path string
	var mainFile string
	var subFiles []string

	generalConvert(&application)

	data, err = style.Data(&application)
	if err != nil {
		return
	}
	path, mainFile, subFiles = style.Files()

	err = compiler.CreateTexFile(data, path, mainFile, subFiles...)
	if err != nil {
		return
	}

	err = compiler.Compile()
	if err != nil {
		return
	}

	//err = compiler.CleanUp()
	return
}

func generalConvert(application *genmodel.Application) {
	util.JoinExperience(application)
}
