package generator

import (
	"errors"
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside"
)

type Style int

const (
	TwoSide_long = Style(iota)
	TwoSide_short
)

func (s Style) Data(application *genmodel.Application) (data interface{}, err error) {
	switch s {
	case TwoSide_long:
		return twoside.Data(application, false)
	case TwoSide_short:
		return twoside.Data(application, true)
	}

	err = errors.New("Style not found.")
	return
}

func (s Style) Files() (path string, mainFile string, subFiles []string) {
	switch s {
	case TwoSide_long, TwoSide_short:
		return twoside.Files()
	}
	return
}
