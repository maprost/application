package util_test

import (
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/assertion"
	"testing"
)

func TestDefaultColor(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.DefaultColor(""), util.DefaultColorValue)
	assert.Equal(util.DefaultColor("              "), util.DefaultColorValue)
	assert.Equal(util.DefaultColor("223355"), "223355")
}

func TestDefaultImage(t *testing.T) {
	assert := assertion.New(t)

	assert.Equal(util.DefaultImage(""), util.NoImagePath)
	assert.Equal(util.DefaultImage("              "), util.NoImagePath)
	assert.Equal(util.DefaultImage("/image"), "/image")
}
