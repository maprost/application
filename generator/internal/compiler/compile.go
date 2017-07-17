package compiler

import (
	"os"
	"text/template"
)

const (
	binaryBase = "application"
	binaryTex  = binaryBase + ".tex"
)

func CreateTexFile(data interface{}, path string, mainFile string, subFiles ...string) (err error) {
	indexFile, err := template.ParseFiles(path + mainFile)
	if err != nil {
		return
	}
	for _, subFile := range subFiles {
		err = addSubTex(indexFile, path, subFile)
		if err != nil {
			return
		}
	}

	// create tex file to compile
	err = createCompilationFile(binaryTex, indexFile, data)
	return
}

func Compile() (err error) {
	// compile
	_, err = stream("pdflatex", binaryTex)
	return
}

func CleanUp() (err error) {
	_, err = stream("rm", binaryBase+".out", binaryBase+".aux", binaryBase+".log", binaryTex)
	return
}

func createCompilationFile(path string, templ *template.Template, data interface{}) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}

	defer file.Close()
	file.Chmod(0644)

	err = templ.Execute(file, data)
	if err != nil {
		return
	}

	return
}

func addSubTex(indexFile *template.Template, path string, subFile string) (err error) {
	subTex, err := template.ParseFiles(path + subFile)
	if err != nil {
		return
	}

	indexFile.AddParseTree(subFile, subTex.Tree)
	return
}
