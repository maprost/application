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
			Pic:         test.ImagePath() + "logan",
			Name:        "Logan van der Bommel",
			Title:       "Good night bringer and tooth painter",
			Nationality: "Canada",
			Email:       "logan@xmen.com",
			Location:    "X-Men Headquarter",
			Phone:       "0123456789",
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
			Skills: []texmodel.Skill{
				{
					Name:   "Cleaning",
					Rating: 4,
				},
				{
					Name:   "Roaring",
					Rating: 5,
				}, {
					Name:   "..",
					Rating: 3,
				}, {
					Name:   "See invisible people with red hair and blue eyes",
					Rating: 4,
				}, {
					Name:   "..",
					Rating: 2,
				},
			},
			OtherSkills: "kill the blob, see the light, have a break, make a bloody marie",
			Language: []texmodel.Language{
				{
					Name:  "German",
					Level: "Native",
				},
				{
					Name:  "English",
					Level: "B2",
				},
			},
			Experience: []texmodel.Experience{
				{
					Company:     "Freelancer",
					Time:        "2000 - 2017",
					Position:    "Peacemaker",
					Description: "Looking for bad guys.",
					Tech:        "MP5",
				},
			},
		},
	}

	err := create(indexData)
	assert.Nil(err)

	err = compile()
	assert.Nil(err)
}
