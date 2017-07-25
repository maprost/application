package internal

import (
	"github.com/maprost/application/generator/genmodel"
	"path"
	"runtime"
)

const (
	techSkill_Go = genmodel.SkillID(iota)
	techSkill_Java
	techSkill_Cpp
	techSkill_Latex
	techSkill_SoftwareDesign
)

const (
	softSkillFriendly = genmodel.SkillID(iota)
)

func Profile() genmodel.Profile {
	_, file, _, _ := runtime.Caller(1)

	return genmodel.Profile{
		FirstName:   "maprost",
		LastName:    "",
		Image:       path.Dir(file) + "/images/cv",
		Birthday:    "...",
		Email:       "...",
		Nationality: "german",
		Websites:    []string{"https://github.com/maprost"},
		Address: genmodel.Address{
			Street:  "...",
			Zip:     "...",
			City:    "...",
			Country: "Germany",
		},
		ProfessionalSkills: map[genmodel.SkillID]genmodel.Skill{
			techSkill_Go:             {Name: "Go", Rating: 8},
			techSkill_Java:           {Name: "Java", Rating: 8},
			techSkill_SoftwareDesign: {Name: "Software Design", Rating: 8},
			techSkill_Cpp:            {Name: "C++", Rating: 5},
			techSkill_Latex:          {Name: "\\LaTeX", Rating: 7},
		},
		SoftSkills: map[genmodel.SkillID]genmodel.Skill{
			softSkillFriendly: {Name: "friendly", Rating: 2},
		},
		Interest: []string{"CleanCode", ""},
		Hobbies:  []string{"Boardgames"},
		Language: []genmodel.Language{
			{
				Name:  "English",
				Level: "B2",
			},
			{
				Name:  "German",
				Level: "Native",
			},
		},
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
