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
	Pub_Calliope = genmodel.ID(iota) + 1
	Pub_Machine
)

const (
	Awa_Fami = genmodel.ID(iota) + 1
	Awa_Mio
)

const (
	Time_1632 = genmodel.ID(iota) + 1
	Time_2432
)

const (
	Money_36for24 = genmodel.ID(iota) + 1
	Money_48for32
	Money_60for40
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
		//LeftSideActionType: []genmodel.LeftSideActionType{
		//	genmodel.Languages,
		//	genmodel.Hobbies,
		//	genmodel.SoftSkills,
		//	genmodel.Interests,
		//	genmodel.TechSkill,
		//	genmodel.TimeAmount,
		//	genmodel.MoneyAmount,
		//},
		TimeAmount: []genmodel.LeftSideAction{
			{Id: Time_1632, Name: lang.TranslationMap{lang.English: "negotiable wish", lang.German: "vehandelbarer Wunsch"}, Min: 16, Max: 32},
			{Id: Time_2432, Name: lang.TranslationMap{lang.English: "negotiable wish", lang.German: "vehandelbarer Wunsch"}, Min: 24, Max: 32},
			//{Name: lang.TranslationMap{lang.English: "negotiable wish 2", lang.German: "vehandelbarer Wunsch 2"}, Min: 6, Max: 32},
		},
		MoneyAmount: []genmodel.LeftSideAction{
			{Id: Money_36for24, Name: lang.TranslationMap{lang.English: "for 24h", lang.German: "für 24h"}, SingleLabel: lang.TranslationMap{lang.English: "36-42T€", lang.German: "36-42T€"}, Min: 36, Max: 42},
			{Id: Money_48for32, Name: lang.TranslationMap{lang.English: "for 32h", lang.German: "für 32h"}, SingleLabel: lang.TranslationMap{lang.English: "48-56T€", lang.German: "48-56T€"}, Min: 48, Max: 56},
			{Id: Money_60for40, Name: lang.TranslationMap{lang.English: "for fulltime", lang.German: "für Vollzeit"}, SingleLabel: lang.TranslationMap{lang.English: "60-70T€", lang.German: "60-70T€"}, Min: 60, Max: 70},
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
			{Name: lang.DefaultTranslation("Blob")},
			//			{Name: lang.DefaultTranslation(`
			//
			//\medskip
			//
			//Blob`)},
			{Name: lang.DefaultTranslation("TDD")},
			{Name: lang.DefaultTranslation("Backend development")},
			{Name: lang.DefaultTranslation("Big Data")},
		},
		Hobbies: []genmodel.LeftSideAction{
			{Name: lang.TranslationMap{lang.English: "Boardgames", lang.German: "Brettspiele"}},
			{Name: lang.TranslationMap{lang.English: "Traveling by bike", lang.German: "Fahrradtouren"}},
			{Name: lang.TranslationMap{lang.English: "aikido", lang.German: "Aikido"}},
		},
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
		GeneralProfileText: lang.TranslationMap{
			lang.English: "I will.",
			lang.German:  "Ich will.",
		},
		//GeneralMyMotivationText: lang.TranslationMap{
		//	lang.English: "GeneralMyMotivationText",
		//	lang.German:  "GeneralMyMotivationText",
		//},
		//CustomMainQuestionTextLabel: lang.TranslationMap{
		//	lang.English: "xxx",
		//	lang.German:  "xxx",
		//},
		//GeneralMainQuestionText: lang.TranslationMap{
		//	lang.English: "GeneralMainQuestionText",
		//	lang.German:  "GeneralMainQuestionText",
		//},
		//CustomMyMotivationTextLabel: lang.TranslationMap{
		//	lang.English: "CustomMyMotivationTextLabel",
		//	lang.German:  "CustomMyMotivationTextLabel",
		//},
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
				Project: lang.TranslationMap{lang.English: "Project 1",
					lang.German: "Projekt 1"},
			},
			{
				Id:          Exp_solar,
				Company:     "Solar Company 123",
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
				Id:            Edu_Uni,
				Graduation:    lang.DefaultTranslation("Diplom Informatiker"),
				Institute:     "German University",
				StartTime:     "Sep. 2005",
				EndTime:       "Oct. 2010",
				Focus:         lang.DefaultTranslation("software engineering, compiler construction, database design"),
				FinalGrade:    lang.DefaultTranslation("1.0"),
				DocumentLinks: []string{"http://google.de"},
				QuitReason:    lang.DefaultTranslation("für den Ausbau der Selbstständigkeit auf eigenen Wunsch beendet"),
			},
			{
				Id:            Edu_Comp,
				Graduation:    lang.DefaultTranslation("Fachinformatiker Systemintegration"),
				Institute:     "Telecommunication Company",
				StartTime:     "Sep. 2002",
				EndTime:       "Jan. 2005",
				Focus:         lang.DefaultTranslation("C++, website development (PHP/mssql), system configuration, router programming"),
				DocumentLinks: []string{"http://google.de", "http://ecosia.de"},
				QuitReason:    lang.DefaultTranslation("für den Ausbau der Selbstständigkeit auf eigenen Wunsch beendet für den Ausbau der Selbstständigkeit auf eigenen Wunsch beendet"),
			},
		},
		Publication: []genmodel.Publication{
			{
				Id:            Pub_Calliope,
				Title:         lang.DefaultTranslation("Das Calliope Buch"),
				Publisher:     lang.DefaultTranslation("XXX"),
				StartTime:     "2005",
				EndTime:       "2005",
				Focus:         lang.DefaultTranslation("XXX"),
				DocumentLinks: []string{"http://google.de"},
			},
			//{
			//	Id:            Pub_Machine,
			//	Graduation:    lang.DefaultTranslation("Fachinformatiker Systemintegration"),
			//	Institute:     "Telecommunication Company",
			//	StartTime:     "Sep. 2002",
			//	EndTime:       "Jan. 2005",
			//	Focus:         lang.DefaultTranslation("C++, website development (PHP/mssql), system configuration, router programming"),
			//	DocumentLinks: []string{"http://google.de", "http://ecosia.de"},
			//},
		},
		Award: []genmodel.Award{
			{
				Id:            Awa_Fami,
				Title:         lang.DefaultTranslation("Beste XY"),
				Institute:     lang.DefaultTranslation("XXX"),
				StartTime:     "Sep. 2005",
				EndTime:       "Oct. 2010",
				Focus:         lang.DefaultTranslation("XXX"),
				DocumentLinks: []string{"http://google.de"},
			},
			//{
			//	Id:            Pub_Machine,
			//	Graduation:    lang.DefaultTranslation("Fachinformatiker Systemintegration"),
			//	Institute:     "Telecommunication Company",
			//	StartTime:     "Sep. 2002",
			//	EndTime:       "Jan. 2005",
			//	Focus:         lang.DefaultTranslation("C++, website development (PHP/mssql), system configuration, router programming"),
			//	DocumentLinks: []string{"http://google.de", "http://ecosia.de"},
			//},
		},
	}
}
