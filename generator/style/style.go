package style

import (
	"errors"
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/onepage"
)

type Style int

const (
	OneSide = Style(iota)
)

func (s Style) Data(application *genmodel.Application) (data interface{}, err error) {
	switch s {
	case OneSide:
		return onepage.Data(application)
	}

	err = errors.New("Style not found.")
	return
}

func (s Style) Files() (path string, mainFile string, subFiles []string) {
	switch s {
	case OneSide:
		return onepage.Files()
	}
	return
}
