package util

import (
	"github.com/maprost/application/generator/internal/image"
	"strings"
)

const (
	DefaultColorValue = "e3593b"
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

func DefaultValue(value string, defaultValue string) string {
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}
	return value
}
