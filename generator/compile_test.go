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
			Location:    "X-Men Headquarter near the Everglades",
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
			OtherSkills: "kill the blob, see the light, have a break, make a bloody marie",
			Interest:    "Datamining",
			SoftSkills:  "Problemsolver",
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
					Description: "Smell and task chocolate. Invented the task bitter sweet.",
					Tech:        "X-Carb3.",
				},
			},
		},
	}

	err := create(indexData)
	assert.Nil(err)

	err = compile()
	assert.Nil(err)
}
