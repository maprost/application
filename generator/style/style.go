package style

import (
	"errors"
	"github.com/maprost/application/generator/genmodel"
	twoside "github.com/maprost/application/generator/internal/style/twoside"
)

type Style int

const (
	OneSide = Style(iota)
	TwoSide
)

func (s Style) Data(app *genmodel.Application, outputPath string) (data interface{}, err error) {
	switch s {
	case OneSide:
		// convert oneSide into twoSide
		app.JobPosition.TwoSideStyle = app.JobPosition.OneSideStyle.TwoSideStyle()
		app.JobPosition.Style = TwoSide
		return twoside.Data(app, outputPath)

	case TwoSide:
		return twoside.Data(app, outputPath)

	}

	err = errors.New("Style not found.")
	return
}

func (s Style) Files() (path string, mainFile string, subFiles []string) {
	switch s {
	case OneSide:
		return twoside.Files()
	case TwoSide:
		return twoside.Files()
	}
	return
}
