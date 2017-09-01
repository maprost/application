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
	techSkill_PostgreSQL
	techSkill_Docker
	techSkill_Git
	techSkill_Python
	techSkill_Bash
	techSkill_HTML5
	techSkill_DBDesign
	techSkill_Testing
)

const (
	softSkill_AnalyticalThinking = genmodel.SkillID(iota)
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
			techSkill_DBDesign:       {Name: "Database Design", Rating: 8},
			techSkill_Testing:        {Name: "Testing", Rating: 9},
			techSkill_Latex:          {Name: "\\LaTeX", Rating: 7},
			techSkill_Cpp:            {Name: "C++", Rating: 5},
			techSkill_IntelliJ:       {Name: "IntelliJ", Rating: 6},
			techSkill_Ubuntu:         {Name: "Ubuntu", Rating: 6},
			techSkill_PostgreSQL:     {Name: "PostgreSQL", Rating: 6},
			techSkill_Docker:         {Name: "Docker", Rating: 5},
			techSkill_Git:            {Name: "Git", Rating: 5},
			techSkill_Python:         {Name: "Python", Rating: 5},
			techSkill_Bash:           {Name: "Bash", Rating: 5},
			techSkill_HTML5:          {Name: "HTML5, JS, CSS", Rating: 4},
		},
		SoftSkills: map[genmodel.SkillID]genmodel.Skill{
			softSkill_AnalyticalThinking: {Name: "Analytical Thinking", Rating: 9},
		},
		Interest: []string{"CleanCode", "Design Pattern", "Microservices", "TDD", "Backend development, Big Data"},
		Hobbies:  []string{"Boardgames", "Traveling by bike"},
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
		GeneralMotivationText: `

		`,
		Experience: []genmodel.Experience{
			{
				Company:     "Fitness Company",
				JobPosition: "Java Backend Developer",
				StartTime:   "Jul. 2014",
				Description: "" +
					"Part of the recruiting team since Oct. 2015.\n" +
					"Product owner of a sub team from Jan. 2016 to Oct. 2016.",
				TechStack: "Java, Guice, GWT, MySQL, Hibernate, RESTful, Json, Docker, Go, \\LaTeX, Ubuntu, Bash, " +
					"Python, Git, Scrum/Kanban, JIRA",
			},
			{
				Company:     "Solar Company",
				JobPosition: "C++ Backend Developer",
				StartTime:   "Mar. 2011",
				EndTime:     "May 2014",
				Description: "Create an XML ",
				TechStack:   "C++, cmake, Boost, PostgreSQL, XML, Bash, Python, OpenSuse, Valgrind, SVN, Mantis",
			},
			{
				Company:     "German University",
				JobPosition: "Tutor",
				StartTime:   "Oct. 2008",
				EndTime:     "Sep. 2010",
				Description: "Teach first and second semester students the basics of programming",
				TechStack:   "Java, object orientation, algorithms, data structures",
			},
		},
		Education: []genmodel.Education{
			{
				Graduation: "Diplom Informatiker",
				Institute:  "German University",
				StartTime:  "Sep. 2005",
				EndTime:    "Oct. 2010",
				Focus:      "software engineering, compiler construction, database design",
				FinalGrade: "1.0",
			},
			{
				Graduation: "Fachinformatiker Systemintegration",
				Institute:  "Telecommunication Company",
				StartTime:  "Sep. 2002",
				EndTime:    "Jan. 2005",
				Focus:      "C++, website development (PHP/mssql), system configuration, router programming",
			},
		},
	}
}
