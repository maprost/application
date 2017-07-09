package generator

import (
	"github.com/maprost/application/generator/internal/image"
	"github.com/maprost/application/generator/internal/test"
	"github.com/maprost/application/generator/internal/texmodel"
	"github.com/maprost/assertion"
	"testing"
)

func TestExample(t *testing.T) {
	assert := assertion.New(t)

	indexData := texmodel.Index{
		CV: texmodel.CV{
			Pic:     test.ImagePath() + "logan",
			Name:    "Logan",
			Title:   "Good Night Bringer",
			Email:   "logan@xmen.com",
			Address: "X-Men Headquarter",
			Phone:   "---",
			Websites: []texmodel.Websites{
				{
					Icon: image.ImagePath() + "github",
					Url:  "https://github.com",
				},
				{
					Icon: image.ImagePath() + "linkedIn",
					Url:  "https://www.linkedin.com",
				},
				{
					Icon: image.ImagePath() + "xing",
					Url:  "https://www.xing.com",
				},
			},
		},
	}

	err := create(indexData)
	assert.Nil(err)

	err = compile()
	assert.Nil(err)
}
