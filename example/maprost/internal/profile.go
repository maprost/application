package internal

import (
	"path"
	"runtime"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
)

const (
	TechSkill_Go = genmodel.ID(iota) + 1
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
	softSkill_AnalyticalThinking = genmodel.ID(iota) + 1
)

const (
	Exp_fitness = genmodel.ID(iota) + 1
	Exp_solar
	Exp_tutor
)

const (
	Edu_Uni = genmodel.ID(iota) + 1
	Edu_Comp
)

const (
	Lang_De = genmodel.ID(iota) + 1
	Lang_En
)

func Profile() genmodel.Profile {
	_, file, _, _ := runtime.Caller(1)

	return genmodel.Profile{
		FirstName: "maprost",
		LastName:  "",
		Image:     path.Dir(file) + "/images/cv",
		Birthday:  "",
		Email:     "N/A",
		Nationality: lang.TranslationMap{
			lang.English: "german",
			lang.German:  "deutsch",
		},
		Phone:    "N/A",
		Websites: []string{"https://github.com/maprost"},
		Address: genmodel.ProfileAddress{
			Street: "",
			Zip:    "",
			City:   lang.DefaultTranslation(""),
			Country: lang.TranslationMap{
				lang.English: "Germany",
				lang.German:  "Deutschland",
			},
		},
		LeftSideActionType: []genmodel.LeftSideActionType{
			genmodel.Languages,
			genmodel.Hobbies,
			genmodel.SoftSkills,
			genmodel.Interests,
			genmodel.TechSkill,
		},
		ProfessionalSkills: []genmodel.LeftSideAction{
			{Id: TechSkill_Go, Name: lang.DefaultTranslation("Go"), Rating: 8},
			{Id: techSkill_Java, Name: lang.DefaultTranslation("Java"), Rating: 9},
			{Id: techSkill_SoftwareDesign, Name: lang.DefaultTranslation("Software Design"), Rating: 9},
			{Id: techSkill_DBDesign, Name: lang.DefaultTranslation("Database Design"), Rating: 8},
			{Id: techSkill_Testing, Name: lang.DefaultTranslation("Testing"), Rating: 9},
			{Id: techSkill_Latex, Name: lang.DefaultTranslation("\\LaTeX"), Rating: 7},
			{Id: techSkill_Cpp, Name: lang.DefaultTranslation("C++"), Rating: 5},
			{Id: techSkill_IntelliJ, Name: lang.DefaultTranslation("IntelliJ"), Rating: 6},
			{Id: techSkill_Ubuntu, Name: lang.DefaultTranslation("Ubuntu"), Rating: 6},
			{Id: techSkill_PostgreSQL, Name: lang.DefaultTranslation("PostgreSQL"), Rating: 6},
			{Id: techSkill_Docker, Name: lang.DefaultTranslation("Docker"), Rating: 5},
			{Id: techSkill_Git, Name: lang.DefaultTranslation("Git"), Rating: 5},
			{Id: techSkill_Python, Name: lang.DefaultTranslation("Python"), Rating: 5},
			{Id: techSkill_Bash, Name: lang.DefaultTranslation("Bash"), Rating: 5},
			{Id: techSkill_HTML5, Name: lang.DefaultTranslation("HTML5, JS, CSS"), Rating: 4},
		},
		SoftSkills: []genmodel.LeftSideAction{
			{Id: softSkill_AnalyticalThinking, Name: lang.TranslationMap{lang.English: "Analytical Thinking", lang.German: "Analytisches Denken"}, Rating: 9},
		},
		CustomInterestLabel: map[lang.Language]string{
			lang.German: "Interessen und Freizeit",
		},
		Interest: []genmodel.LeftSideAction{
			{Name: lang.DefaultTranslation("CleanCode")},
			{Name: lang.DefaultTranslation("Design Pattern")},
			{Name: lang.DefaultTranslation("Microservices")},
			{Name: lang.DefaultTranslation(`

\medskip

Blob`)},
			{Name: lang.DefaultTranslation("TDD")},
			{Name: lang.DefaultTranslation("Backend development")},
			{Name: lang.DefaultTranslation("Big Data")},
		},
		//Hobbies: []lang.TranslationMap{
		//	{lang.English: "Boardgames", lang.German: "Brettspiele"},
		//	{lang.English: "Traveling by bike", lang.German: "Fahrradtouren"},
		//	{lang.English: "aikido", lang.German: "Aikido"},
		//},
		Language: []genmodel.Language{
			{
				Id:    Lang_De,
				Name:  lang.TranslationMap{lang.English: "German", lang.German: "Deutsch"},
				Level: lang.TranslationMap{lang.English: "Native", lang.German: "Muttersprache"},
			},
			{
				Id:    Lang_En,
				Name:  lang.TranslationMap{lang.English: "English", lang.German: "Englisch"},
				Level: lang.DefaultTranslation("B2"),
			},
		},
		GeneralMotivationText: lang.TranslationMap{
			lang.English: "I will.",
			lang.German:  "Ich will.",
		},
		Experience: []genmodel.Experience{
			{
				Id:          Exp_fitness,
				Company:     "Fitness Company",
				JobPosition: lang.DefaultTranslation("Java Backend Developer"),
				StartTime:   "Jul. 2014",
				Description: lang.TranslationMap{lang.English: "" +
					"Part of the recruiting team since Oct. 2015.\n" +
					"Product owner of a sub team from Jan. 2016 to Oct. 2016.",
					lang.German: "Recruiter, ProductOwner"},
				TechStack: lang.DefaultTranslation("Java, Guice, GWT, MySQL, Hibernate, RESTful, Json, Docker, Go, \\LaTeX, Ubuntu, Bash, " +
					"Python, Git, Scrum/Kanban, JIRA"),
				DocumentLinks: []string{"http://google.de"},
			},
			{
				Id:          Exp_solar,
				Company:     "Solar Company",
				JobPosition: lang.DefaultTranslation("C++ Backend Developer"),
				StartTime:   "Mar. 2011",
				EndTime:     "May 2014",
				Description: lang.TranslationMap{lang.English: "",
					lang.German: ""},
				TechStack:  lang.DefaultTranslation("C++, cmake, Boost, PostgreSQL, XML, Bash, Python, OpenSuse, Valgrind, SVN, Mantis"),
				TechStack2: lang.DefaultTranslation("C++, PostgreSQL"),
			},
			{
				Id:          Exp_tutor,
				Company:     "German University",
				JobPosition: lang.DefaultTranslation("Tutor"),
				StartTime:   "Oct. 2008",
				EndTime:     "Sep. 2010",
				Description: lang.TranslationMap{lang.English: "Teach first and second semester students the basics of programming",
					lang.German: ""},
				TechStack: lang.DefaultTranslation("Java, object orientation, algorithms, data structures"),
			},
		},
		Education: []genmodel.Education{
			{
				Id:         Edu_Uni,
				Graduation: lang.DefaultTranslation("Diplom Informatiker"),
				Institute:  "German University",
				StartTime:  "Sep. 2005",
				EndTime:    "Oct. 2010",
				Focus:      lang.DefaultTranslation("software engineering, compiler construction, database design"),
				FinalGrade: lang.DefaultTranslation("1.0"),
			},
			{
				Id:         Edu_Comp,
				Graduation: lang.DefaultTranslation("Fachinformatiker Systemintegration"),
				Institute:  "Telecommunication Company",
				StartTime:  "Sep. 2002",
				EndTime:    "Jan. 2005",
				Focus:      lang.DefaultTranslation("C++, website development (PHP/mssql), system configuration, router programming"),
			},
		},
	}
}
