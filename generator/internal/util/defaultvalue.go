package util

import (
	"strings"

	"github.com/maprost/application/generator/internal/image"
)

const (
	DefaultColorValue     = "e3593b"
	DefaultSideColorValue = "e7e7e7"
)

var (
	NoImagePath = image.ImagePath() + "noimage"
)

func DefaultColor(color string) string {
	return DefaultValue(color, DefaultColorValue)
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
