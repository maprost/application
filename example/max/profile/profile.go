package profile

import (
	"github.com/maprost/application/generator/genmodel"
)

const (
	TechSkillHiding = genmodel.ID(iota)
	TechSkillSneaking
	TechSkillWrapping
	TechSkillClimbing
)

const (
	SoftSkillFriendly = genmodel.ID(iota)
)

func New() genmodel.Profile {
	//_, file, _, _ := runtime.Caller(1)

	return genmodel.Profile{
		//FirstName:   "Mad",
		//LastName:    "Max",
		//Image:       path.Dir(file) + "/images/cv",
		//Birthday:    "1980, California",
		//Email:       "mad@max.com",
		//Nationality: "american",
		//Websites:    []string{"https://www.linkedin.com", "www.myspace.com"},
		//Address: genmodel.Address{
		//	Street:  "NewStr. 56",
		//	Zip:     "12345",
		//	City:    "LA",
		//	Country: "USA",
		//},
		//ProfessionalSkills: map[genmodel.ID]genmodel.Skill{
		//	TechSkillHiding:   {Name: "hiding", Rating: 3},
		//	TechSkillSneaking: {Name: "sneaking", Rating: 5},
		//	TechSkillWrapping: {Name: "wrapping", Rating: 3},
		//	TechSkillClimbing: {Name: "climbing", Rating: 4},
		//},
		//SoftSkills: map[genmodel.ID]genmodel.Skill{
		//	SoftSkillFriendly: {Name: "friendly", Rating: 2},
		//},
		//Interest: []string{"Survival training", "Water"},
		//Hobbies:  []string{"cars"},
		//Language: []genmodel.Language{
		//	{
		//		Name:  "English",
		//		Level: "C2",
		//	},
		//	{
		//		Name:  "French",
		//		Level: "A2",
		//	},
		//},
		//Experience: []genmodel.Experience{
		//	{
		//		Company:     "Lord Drug",
		//		JobPosition: "Joker",
		//		StartTime:   "Jun. 2014",
		//		EndTime:     "Jul 2026",
		//		Description: "Bring some fun to the people in my environment and let them laugh.",
		//		TechStack:   "MP5, Knifes",
		//	},
		//},
		//Education: []genmodel.Education{
		//	{
		//		Institute: "LA Institute of Nowhere",
		//		StartTime: "Aug 2000",
		//		EndTime:   "Sep 2000",
		//		Focus:     "Learn how to survive in a difficult area.",
		//	},
		//},
	}
}
