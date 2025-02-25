package twoside

import (
	"path/filepath"
	"runtime"

	"github.com/maprost/application/generator/genmodel"
)

func Data(application *genmodel.Application, outputPath string) (data interface{}, err error) {
	index, err := initData(application, outputPath)
	if err != nil {
		return
	}

	data = index
	return
}

func Files() (path string, mainFile string, subFiles []string) {
	path = templatePath()
	mainFile = "index.tex"
	subFiles = []string{}
	return
}

func templatePath() string {
	return rootPath() + "/template/"
}

func rootPath() string {
	_, curFile, _, _ := runtime.Caller(0)
	return filepath.Dir(curFile)
}
