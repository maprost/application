package onepage

import (
	"path/filepath"
	"runtime"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
)

func Data(application *genmodel.Application, lang lang.Language) (data interface{}, err error) {
	index, err := initData(application, lang)
	if err != nil {
		return
	}

	data = index
	return
}

func Files() (path string, mainFile string, subFiles []string) {
	path = rootPath() + "/template/"
	mainFile = "index.tex"
	subFiles = []string{}
	return
}

func rootPath() string {
	_, curFile, _, _ := runtime.Caller(0)
	return filepath.Dir(curFile)
}
