package twoside

import (
	"path/filepath"
	"runtime"

	"github.com/maprost/application/generator/genmodel"
)

func Data(application *genmodel.Application) (data interface{}, err error) {
	index, err := initIndex(application)

	index.FirstPage, err = createFirstPageData(application)
	if err != nil {
		return
	}

	index.CoverLetter, err = createCoverLetterData(application)
	if err != nil {
		return
	}

	index.CV, err = createCVData(application)
	if err != nil {
		return
	}

	data = index
	return
}

func Files() (path string, mainFile string, subFiles []string) {
	path = rootPath() + "/template/"
	mainFile = "index.tex"
	subFiles = []string{"firstpage.tex", "coverletter.tex", "cv.tex"}
	return
}

func rootPath() string {
	_, curFile, _, _ := runtime.Caller(0)
	return filepath.Dir(curFile)
}
