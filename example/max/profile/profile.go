package profile

import (
	"github.com/maprost/application/generator/genmodel"
	"path"
	"runtime"
)

const (
	TechSkillHiding = genmodel.SkillID(iota)
	TechSkillSneaking
	TechSkillWrapping
	TechSkillClimbing
)

const (
	SoftSkillFriendly = genmodel.SkillID(iota)
)

func New() genmodel.Profile {
	_, file, _, _ := runtime.Caller(1)

	return genmodel.Profile{
		FirstName:   "Mad",
		LastName:    "Max",
		Image:       path.Dir(file) + "/images/cv",
		Birthday:    "1980, California",
		Email:       "mad@max.com",
		Nationality: "american",
		Website:     []string{"https://www.linkedin.com"},
		Address: genmodel.Address{
			Street:  "NewStr. 56",
			Zip:     "12345",
			City:    "LA",
			Country: "USA",
		},
		TechnicalSkills: map[genmodel.SkillID]genmodel.Skill{
			TechSkillHiding:   {Name: "hiding", Rating: 3},
			TechSkillSneaking: {Name: "sneaking", Rating: 5},
			TechSkillWrapping: {Name: "wrapping", Rating: 3},
			TechSkillClimbing: {Name: "climbing", Rating: 4},
		},
		SoftSkills: map[genmodel.SkillID]genmodel.Skill{
			SoftSkillFriendly: {Name: "friendly", Rating: 2},
		},
		Interest: []string{"Survival training", "Water"},
		Hobbies:  []string{"cars"},
		Experience: []genmodel.Experience{
			{
				Company:     "Lord Drug",
				JobPosition: "Joker",
				StartTime:   "Jun. 2014",
				EndTime:     "Jul 2026",
				Description: "Bring some fun to the people in my environment and let them laugh.",
				TechStack:   "MP5, Knifes",
			},
		},
		Education: []genmodel.Education{
			{
				School:      "LA School of Nowhere",
				StartTime:   "Aug 2000",
				EndTime:     "Sep 2000",
				Description: "Learn how to survive in a difficult area.",
			},
		},
	}
}
