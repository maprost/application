package generator

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/compiler"
)

func Build(profile genmodel.Profile, style Style) (err error) {
	var data interface{}
	var path string
	var mainFile string
	var subFiles []string

	data, err = style.Data(profile)
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

	err = compiler.CleanUp()
	return
}
