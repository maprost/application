package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"path/filepath"
	"runtime"
)

func Data(profile genmodel.Profile) (data interface{}, err error) {

	return
}

func Files() (path string, mainFile string, subFiles []string) {
	path = rootPath() + "/template/"
	mainFile = "index.tex"
	subFiles = []string{"cv.tex"}
	return
}

func rootPath() string {
	_, curFile, _, _ := runtime.Caller(0)
	return filepath.Dir(curFile)
}
