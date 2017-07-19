package util

import (
	"github.com/maprost/application/generator/internal/image"
	"strings"
)

func DefaultColor(color string) string {
	return defaultValue(color, "e3593b")
}

func DefaultImage(path string) string {
	return defaultValue(path, image.ImagePath()+"noimage")
}

func defaultValue(value string, defaultValue string) string {
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}
	return value
}
