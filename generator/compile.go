package generator

import (
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/convert"
	"github.com/maprost/application/generator/internal/shell"
)

var root = rootPath()

func Build(profile genmodel.Profile) (err error) {
	err = convert.Convert(&profile)
	if err != nil {
		return
	}

	// create templates
	index, err := template.ParseFiles(root + "/internal/template/index.tex")
	if err != nil {
		return
	}
	addSubTex(index, "cv.tex")

	// create tex file to compile
	compilationFileBase := "application"
	compilationFile := compilationFileBase + ".tex"
	err = createCompilationFile(compilationFile, index, profile)
	if err != nil {
		return
	}

	// compile file
	_, err = shell.Stream("pdflatex", compilationFile)
	if err != nil {
		return
	}

	// clean up
	_, err = shell.Stream("rm", compilationFileBase+".aux", compilationFileBase+".log", compilationFile)
	return
}

func rootPath() string {
	_, curFile, _, _ := runtime.Caller(0)
	return filepath.Dir(curFile)
}

func createCompilationFile(path string, templ *template.Template, profile genmodel.Profile) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}

	defer file.Close()
	file.Chmod(0644)

	err = templ.Execute(file, profile)
	if err != nil {
		return
	}

	return
}

func addSubTex(main *template.Template, file string) (err error) {
	subTex, err := template.ParseFiles(root + "/internal/template/" + file)
	if err != nil {
		return
	}

	main.AddParseTree(file, subTex.Tree)
	return
}
