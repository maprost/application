package compiler

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/util"
)

func CreateTexFile(outputPath string, file string, data interface{}, path string, mainFile string, subFiles ...string) (err error) {
	indexPath := path + mainFile
	//c, err := os.ReadFile(indexPath)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(string(c))

	indexFile, err := template.ParseFiles(indexPath)
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
	err = createCompilationFile(outputPath, texFile("", file), indexFile, data)
	return
}

func Compile(outputPath string, file string) (err error) {
	// compile
	_, err = stream("pdflatex", "-output-directory="+outputPath, texFile(outputPath, file))
	return
}

func CleanUp(outputPath string, file string) (err error) {
	_, err = stream("rm",
		concat(outputPath, file, "aux"),
		concat(outputPath, file, "log"),
		concat(outputPath, file, "out"),
		//texFile(outputPath, file),
	)
	return
}

func createCompilationFile(path string, fileName string, templ *template.Template, data interface{}) (err error) {
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}

	file, err := os.Create(path + "/" + fileName)
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

func texFile(path string, file string) string {
	return concat(path, file, "tex")
}

func concat(path string, file string, end string) string {
	return util.JoinStrings(path, "/", file+"."+end)
}

func CompileSubTex(path string, mainFile string, data interface{}) (string, error) {
	subPath := path + mainFile
	fmt.Println("----------------------------------------------")
	fmt.Println("--- CompileSubTex: ", subPath)
	subFile, err := template.ParseFiles(subPath)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = subFile.Execute(&b, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func GenerateOutput(outputPath string, app *genmodel.Application) (path string, file string) {
	path = app.JobPosition.OutputPath
	if path == "" {
		path = outputPath
	}

	file = app.JobPosition.FileName
	if file == "" {
		file = "application"
	}

	return path, file
}
