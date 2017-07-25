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
	techSkill_IntelliJ
	techSkill_Ubuntu
	techSkill_PostgresSQL
	techSkill_Docker
	techSkill_Git
	techSkill_Docker
	techSkill_Docker
	techSkill_Docker
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
		Birthday:    "",
		Email:       "N/A",
		Nationality: "german",
		Phone:       "N/A",
		Websites:    []string{"https://github.com/maprost"},
		Address: genmodel.Address{
			Street:  "",
			Zip:     "",
			City:    "",
			Country: "Germany",
		},
		ProfessionalSkills: map[genmodel.SkillID]genmodel.Skill{
			techSkill_Go:             {Name: "Go", Rating: 8},
			techSkill_Java:           {Name: "Java", Rating: 9},
			techSkill_SoftwareDesign: {Name: "Software Design", Rating: 9},
			techSkill_Cpp:            {Name: "C++", Rating: 5},
			techSkill_Latex:          {Name: "\\LaTeX", Rating: 7},
			techSkill_IntelliJ:       {Name: "IntelliJ", Rating: 6},
			techSkill_Ubuntu:         {Name: "Ubuntu", Rating: 6},
			techSkill_PostgresSQL:    {Name: "Postgres", Rating: 6},
			techSkill_Docker:         {Name: "Docker", Rating: 5},
			techSkill_Git:            {Name: "Git", Rating: 5},
		},
		SoftSkills: map[genmodel.SkillID]genmodel.Skill{
			softSkillFriendly: {Name: "friendly", Rating: 9},
		},
		Interest: []string{"CleanCode", "TDD", "Design Pattern"},
		Hobbies:  []string{"Boardgames"},
		Language: []genmodel.Language{
			{
				Name:  "German",
				Level: "Native",
			},
			{
				Name:  "English",
				Level: "B2",
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
