package generator

import (
	"errors"
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside"
)

type Style int

const (
	TwoSide = Style(iota)
)

func (s Style) Data(application *genmodel.Application) (data interface{}, err error) {
	switch s {
	case TwoSide:
		return twoside.Data(application)
	}

	err = errors.New("Style not found.")
	return
}

func (s Style) Files() (path string, mainFile string, subFiles []string) {
	switch s {
	case TwoSide:
		return twoside.Files()
	}
	return
}
