package twoside_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/image"
	"github.com/maprost/application/generator/internal/style/twoside"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
	"github.com/maprost/application/generator/internal/test"
	"github.com/maprost/application/generator/internal/util"
)

func TestData_emptyInput(t *testing.T) {
	assertApplication(t, genmodel.Application{}, texmodel.Index{})
}

func assertApplication(t *testing.T, application genmodel.Application, expected texmodel.Index) {
	assert := assertion.New(t)

	data, err := twoside.Data(&application, expected.ShortVersion)
	assert.Nil(err)

	texmodelData, ok := data.(texmodel.Index)
	assert.True(ok)

	// modify (set default values) expected:
	if expected.MainColor == "" {
		expected.MainColor = util.DefaultColorValue
	}

	if !expected.ShortVersion && expected.FirstPage.Image == "" {
		expected.FirstPage.Image = util.NoImagePath
	}

	if expected.CV.Image == "" {
		expected.CV.Image = util.NoImagePath
	}

	assert.Equal(texmodelData, expected)
}

func _TestData_contactData_blob(t *testing.T) {
	assertApplication(t,
		genmodel.Application{},
		texmodel.Index{
			MainColor: "e3593b",
			CV: texmodel.CV{
				Image:       test.ImagePath() + "logan",
				Name:        "Logan van der Bommel",
				Title:       "Good night bringer and tooth painter",
				Nationality: "Canada",
				Email:       "logan@xmen.com",
				Location:    "X-Men Headquarter near the Everglades",
				Phone:       "0123456789",
				Websites: []texmodel.Website{
					{
						Icon: image.ImagePath() + "github",
						Url:  "https://github.com",
					},
					{
						Icon: image.ImagePath() + "linkedIn",
						Url:  "https://linkedin.com",
					},
					{
						Icon: image.ImagePath() + "xing",
						Url:  "https://www.xing.com",
					},
				},
				ProfSkills: []texmodel.Skill{
					{
						Name:   "Cleaning",
						Rating: 7,
					},
					{
						Name:   "Roaring",
						Rating: 9,
					}, {
						Name:   "Sneaking",
						Rating: 2,
					}, {
						Name:   "See invisible people with red hair",
						Rating: 5,
					}, {
						Name:   "Climbing",
						Rating: 7,
					},
				},
				OtherProfSkills: "kill the blob, see the light, have a break, make a bloody marie",
				Interest:        "Datamining",
				SoftSkills:      "Problemsolver",
				Language: []texmodel.Language{
					{
						Name:  "English",
						Level: "Native",
					},
					{
						Name:  "French",
						Level: "C1",
					},
					{
						Name:  "Portuguese",
						Level: "A1",
					},
					{
						Name:  "Hiligaynon/ Ilonggo ",
						Level: "B1+",
					},
				},
				Hobbies: "Skat, Bicycle",
				Experience: []texmodel.Experience{
					{
						Company:     "Freelancer",
						Time:        "2000 - 2017",
						Position:    "Peacemaker",
						Description: "Looking for bad guys.",
						Tech:        "MP5",
					},
					{
						Company:     "Hack'n'Slay",
						Time:        "1995 - 1999",
						Position:    "Senior Tester",
						Description: "Looking for bad guys.",
						Tech:        "MP5",
					},
					{
						Company:     "Freelancer",
						Time:        "1994 - 1996",
						Position:    "Chocolate taste definer",
						Description: "Smell and task chocolate. Invented the taste bitter sweet.",
						Tech:        "X-Carb3.",
					},
				},
				Education: []texmodel.Education{
					{
						Institute: "MIT",
						Time:      "1899-1903",
						Focus:     "Learn a lot and many more and finish my study with 1.0.",
					},
				},
			},
		})
}
