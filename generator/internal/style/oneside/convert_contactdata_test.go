package oneside_test

import (
	"testing"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/onepage/texmodel"
	"github.com/maprost/application/generator/internal/test"
	"github.com/maprost/application/generator/internal/util"
)

func TestData_contactData_longVersion(t *testing.T) {
	assertApplication(t,
		genmodel.Application{
			Profile: genmodel.Profile{
				Image:       test.ImagePath() + "logan",
				FirstName:   "Logan",
				LastName:    "Bommel",
				Nationality: "Canada",
				Birthday:    "1923, Canada",
				Email:       "logan@xmen.com",
				Address: genmodel.Address{
					Street:  "BommelStr. 44",
					City:    "LA",
					Zip:     "1223",
					Country: "USA",
				},
				Phone: "0123456789",
				Websites: []string{
					"https://github.com", "https://linkedin.com", "https://www.xing.com", "https://www.mywebsite.com",
				},
			},
			JobPosition: genmodel.JobPosition{
				Title: "Dreamer",
			},
		},
		texmodel.Index{
			Image:       test.ImagePath() + "logan",
			Name:        "Logan Bommel",
			Title:       "Dreamer",
			Nationality: "Canada",
			Email:       "logan@xmen.com",
			Location:    "LA, USA",
			Phone:       "0123456789",
			Websites: []texmodel.Website{
				{
					Icon: util.GithubIconPath,
					Url:  "https://github.com",
				},
				{
					Icon: util.LinkedinIconPath,
					Url:  "https://linkedin.com",
				},
				{
					Icon: util.XingIconPath,
					Url:  "https://www.xing.com",
				},
				{
					Icon: util.WebsiteIconPath,
					Url:  "https://www.mywebsite.com",
				},
			},
		})
}

func TestData_contactData_shortVersion(t *testing.T) {
	assertApplication(t,
		genmodel.Application{
			Profile: genmodel.Profile{
				Image:       test.ImagePath() + "logan",
				FirstName:   "Logan",
				LastName:    "Bommel",
				Nationality: "Canada",
				Birthday:    "1923, Canada",
				Email:       "logan@xmen.com",
				Address: genmodel.Address{
					Street:  "BommelStr. 44",
					City:    "LA",
					Zip:     "1223",
					Country: "USA",
				},
				Phone: "0123456789",
				Websites: []string{
					"https://github.com", "https://linkedin.com", "https://www.xing.com", "https://www.mywebsite.com",
				},
			},
			JobPosition: genmodel.JobPosition{
				Title: "Dreamer",
			},
		},
		texmodel.Index{
			Image:       test.ImagePath() + "logan",
			Name:        "Logan Bommel",
			Title:       "Dreamer",
			Nationality: "Canada",
			Email:       "logan@xmen.com",
			Location:    "LA, USA",
			Phone:       "0123456789",
			Websites: []texmodel.Website{
				{
					Icon: util.GithubIconPath,
					Url:  "https://github.com",
				},
				{
					Icon: util.LinkedinIconPath,
					Url:  "https://linkedin.com",
				},
				{
					Icon: util.XingIconPath,
					Url:  "https://www.xing.com",
				},
				{
					Icon: util.WebsiteIconPath,
					Url:  "https://www.mywebsite.com",
				},
			},
		})
}
