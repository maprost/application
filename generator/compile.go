package generator

import (
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/convert"
	"github.com/maprost/application/generator/internal/shell"
	"github.com/maprost/application/generator/internal/texmodel"
)

const (
	binaryBase = "application"
	binaryTex  = binaryBase + ".tex"
)

var root = rootPath()

func Build(profile genmodel.Profile) (err error) {
	indexData, err := prepare(&profile)
	if err != nil {
		return
	}

	// create templates
	err = create(indexData)
	if err != nil {
		return
	}

	// compile file
	err = compile()
	if err != nil {
		return
	}

	// remove .aux/.log/.tex file
	cleanUp()
	return
}

func prepare(profile *genmodel.Profile) (texmodel.Index, error) {
	return convert.Convert(profile)
}

func create(indexData texmodel.Index) (err error) {
	indexFile, err := template.ParseFiles(root + "/internal/template/index.tex")
	if err != nil {
		return
	}
	addSubTex(indexFile, "cv.tex")

	// create tex file to compile

	err = createCompilationFile(binaryTex, indexFile, indexData)
	return
}

func compile() (err error) {
	_, err = shell.Stream("pdflatex", binaryTex)
	return
}

func cleanUp() (err error) {
	// clean up
	_, err = shell.Stream("rm", binaryBase+".aux", binaryBase+".log", binaryTex)
	return
}

func rootPath() string {
	_, curFile, _, _ := runtime.Caller(0)
	return filepath.Dir(curFile)
}

func createCompilationFile(path string, templ *template.Template, indexData texmodel.Index) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}

	defer file.Close()
	file.Chmod(0644)

	err = templ.Execute(file, indexData)
	if err != nil {
		return
	}

	return
}

func addSubTex(indexFile *template.Template, file string) (err error) {
	subTex, err := template.ParseFiles(root + "/internal/template/" + file)
	if err != nil {
		return
	}

	indexFile.AddParseTree(file, subTex.Tree)
	return
}
