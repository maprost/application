package twoside

import (
	"path/filepath"
	"runtime"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
)

func Data(application *genmodel.Application) (data interface{}, err error) {
	var index texmodel.Index

	index.FirstPage, err = convertFirstPageData(application)
	if err != nil {
		return
	}

	index.CoverLetter, err = convertCoverLetterData(application)
	if err != nil {
		return
	}

	index.CV, err = convertCVData(application)
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
