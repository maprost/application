package util

import (
	"github.com/maprost/application/generator/genmodel"
	"strings"

	"github.com/maprost/application/generator/internal/image"
)

const (
	DefaultColorValue        = "e3593b"
	DefaultMainColorValue    = "e7e7e7"
	DefaultSideColorValue    = "e7e7e7"
	DefaultShadowColorValue  = "cccccc"
	DefaultScaleBgColorValue = "ffffff"
)

var (
	NoImagePath = image.ImagePath() + "noimage"
)

func DefaultColor(color string) string {
	return DefaultValue(color, DefaultColorValue)
}

func GetSpecificDefaultColor(color string, defaultValue string) string {
	return DefaultValue(color, defaultValue)
}

func DefaultImage(path string) string {
	return DefaultValue(path, NoImagePath)
}

func DefaultStringArray(arr1 []string, arr2 []string) []string {
	if len(arr1) > 0 {
		return arr1
	}
	return arr2
}

func DefaultValue(value string, defaultValue string) string {
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}
	return value
}

func DefaultLeftSideList(value []genmodel.LeftSideAction, defaultValue []genmodel.LeftSideAction) []genmodel.LeftSideAction {
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
