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
	Awa_BSS
)

const (
	Time_1632 = genmodel.ID(iota) + 1
	Time_2432
	Time_20full
	Time_HO2d
	Time_HO16h
	Time_HomeofficeFull
)

const (
	Money_36for24 = genmodel.ID(iota) + 1
	Money_48for32
	Money_60for40
)

const (
	Lang_De = genmodel.ID(iota) + 1
	Lang_DeInvisible
	Lang_En
)

const (
	Doc_CV   = "generator/internal/files/maprost.pdf"
	Doc_Plan = "generator/internal/files/Speiseplan-Kita-Krippe.pdf"
)

func Profile() genmodel.Profile {
	_, file, _, _ := runtime.Caller(1)
	path := path.Dir(file)

	return genmodel.Profile{
		FirstName: "My First",
		LastName:  "My Last",
		Image:     path + "/images/cv",
		Birthday:  "",
		Email:     "mat@test.de",
		Nationality: lang.TranslationMap{
			lang.English: "german",
			lang.German:  "deutsch",
		},
		NationalityMig: lang.TranslationMap{
			lang.English: "german, 2",
			lang.German:  "deutsch, 2",
		},
		Phone:    "N/A",
		Websites: []string{"https://test.com/maprost", "https://github.com/maprost", "https://www.xing.com", "https://www.linkedin.com"},
		Address: genmodel.ProfileAddress{
			Street: "Meine Straße",
			Zip:    "Meine Zip",
			City:   lang.DefaultTranslation("Meine Stadt"),
			Country: lang.TranslationMap{
				lang.English: "Germany",
				lang.German:  "Deutschland",
			},
		},
		SignPath: path + "/images/sig_test",
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
			{Id: Time_20full, Name: lang.TranslationMap{lang.English: "open for part time", lang.German: "offen für Teilzeit"}, Min: 20, Max: 40},
			{Id: Time_HO2d, Name: lang.TranslationMap{lang.English: "open for part time", lang.German: "je nach Aufgabe und Absprache anteilige Homeoffice-Möglichkeit"}, Min: 0, Max: 2, Full: 5, CurrencyEnding: "d"},
			{Id: Time_HO16h, Name: lang.TranslationMap{lang.English: "open for part time", lang.German: "je nach Aufgabe und Absprache anteilige Homeoffice-Möglichkeit"}, Min: 0, Max: 16},
			{Id: Time_HomeofficeFull, Name: lang.TranslationMap{lang.English: "desired homeoffice", lang.German: "gewünschtes Homeoffice"}, Min: 32, Max: 40, SingleLabel: lang.TranslationMap{lang.English: "80-100\\% \\\\\\ HO", lang.German: "80-100\\% \\\\\\ HO"}},
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
				Id:             Lang_DeInvisible,
				Name:           lang.TranslationMap{lang.English: "German2", lang.German: "Deutsch2"},
				Level:          lang.TranslationMap{lang.English: "Native", lang.German: "Muttersprache"},
				MustBeSelected: true,
			},
			{
				Id:             Lang_En,
				Name:           lang.TranslationMap{lang.English: "English", lang.German: "Englisch"},
				Level:          lang.DefaultTranslation("B2"),
				ShowRefExplain: true,
				DocumentLinks:  []string{"http://google.de", "http://google.de"},
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
				Id:             Exp_fitness,
				MustBeSelected: true,
				Company:        "Fitness Company",
				Languages: []genmodel.RunningLanguage{
					{
						Name: map[lang.Language]string{
							lang.German:  "Firmensprache",
							lang.English: "Company language",
						},
						Language: lang.English,
					},
					{
						Name: map[lang.Language]string{
							lang.German:  "Kaffesprache",
							lang.English: "Coffee language",
						},
						Language: lang.German,
					},
				},
				JobPositionFirstLine: lang.DefaultTranslation("Test erste Zeile"),
				JobPosition:          lang.DefaultTranslation("Java Backend Developer"),
				StartTime:            "Jul. 2014",
				Description: lang.TranslationMap{
					lang.German: "Recruiter, ProductOwner",
					lang.English: "" +
						"Part of the recruiting team since Oct. 2015.\n" +
						"Product owner of a sub team from Jan. 2016 to Oct. 2016.",
				},
				TechStack: lang.DefaultTranslation("Java, Guice, GWT, MySQL, Hibernate, RESTful, Json, Docker, Go, \\LaTeX, Ubuntu, Bash, " +
					"Python, Git, Scrum/Kanban, JIRA"),
				TechStack2:    lang.DefaultTranslation("xxx"),
				DocumentLinks: []string{"http://google.de"},
				Project: lang.TranslationMap{
					lang.English: "Project 1",
					lang.German:  "Projekt 1",
				},
				Project2: lang.TranslationMap{
					lang.English: "Project 2",
					lang.German:  "Projekt 2",
				},
			},
			{
				Id:                   Exp_solar,
				Company:              "Solar Company",
				JobPositionFirstLine: lang.DefaultTranslation("Software Entwicklerin,"),
				JobPosition:          lang.DefaultTranslation("Projektleiterin"),
				StartTime:            "Mar. 2011",
				EndTime:              "May 2014",
				Description: lang.TranslationMap{lang.English: "",
					lang.German: ""},
				TechStack:  lang.DefaultTranslation("C++, cmake, Boost, PostgreSQL, XML, Bash, Python, OpenSuse, Valgrind, SVN, Mantis"),
				TechStack2: lang.DefaultTranslation("C++, PostgreSQL"),
				Role:       lang.DefaultTranslation("Meine Rollen"),
				QuitReason: lang.DefaultTranslation("befristeter Vertrag"),
			},
			{
				Id:          Exp_tutor,
				Company:     "German University",
				JobPosition: lang.DefaultTranslation("Tutor"),
				StartTime:   "Oct. 2008",
				EndTime:     "Sep. 2010",
				Description: lang.TranslationMap{lang.English: "Teach first and second semester students the basics of programming",
					lang.German: ""},
				TechStack:  lang.DefaultTranslation("Java, object orientation, algorithms, data structures"),
				QuitReason: lang.DefaultTranslation("befristeter Vertrag"),
			},
		},
		Education: []genmodel.Education{
			{
				Id:                  Edu_Uni,
				GraduationFirstLine: lang.DefaultTranslation("Diplom Informatiker First Line"),
				Graduation:          lang.DefaultTranslation("Diplom Informatiker"),
				Institute:           "German University",
				StartTime:           "Sep. 2005",
				EndTime:             "Oct. 2010",
				Focus:               lang.DefaultTranslation("software engineering, compiler construction, database design"),
				FinalGrade:          lang.DefaultTranslation("1.0"),
				DocumentLinks:       []string{"http://google.de"},
				QuitReason:          lang.DefaultTranslation("für den Ausbau der Selbstständigkeit auf eigenen Wunsch beendet"),
			},
			{
				Id:            Edu_Comp,
				Graduation:    lang.DefaultTranslation("Fachinformatiker Systemintegration"),
				Institute:     "Telecommunication Company",
				StartTime:     "Sep. 2002",
				EndTime:       "Jan. 2005",
				Focus:         lang.DefaultTranslation("C++, website development (PHP/mssql), system configuration, router programming"),
				DocumentLinks: []string{"http://ecosia.de"},
				ExternalLinks: []string{"http://google.de", "http://ecosia.de"},
				QuitReason:    lang.DefaultTranslation("für den Ausbau der Selbstständigkeit auf eigenen Wunsch beendet"),
			},
		},
		Publication: []genmodel.Publication{
			{
				Id:              Pub_Calliope,
				Title:           lang.DefaultTranslation("Das Buch"),
				Publisher:       lang.DefaultTranslation("XXX"),
				Date:            "2005",
				CoverImage:      path + "/images/cv",
				Content:         lang.DefaultTranslation("XXX"),
				ContentShortLsa: lang.DefaultTranslation("short"),
				ExternalLinks:   []string{"http://google.de"},
			},
			{
				Id:              Pub_Machine,
				Title:           lang.DefaultTranslation("Buch 1"),
				Publisher:       lang.DefaultTranslation("Verlag 1"),
				Date:            "2019",
				CoverImage:      "generator/internal/image/noimage.png",
				Content:         lang.DefaultTranslation("Bschreibung 1"),
				ContentShortLsa: lang.DefaultTranslation("short"),
				ExternalLinks:   []string{"http://google.de", "http://google.de"},
			},
			{
				Id: Pub_Calliope,
				Title: lang.TranslationMap{
					lang.German:  `Buch 2`,
					lang.English: `Book 2`,
				},
				Publisher:  lang.DefaultTranslation("Verlag 2"),
				Date:       "2017",
				CoverImage: "generator/internal/image/noimage.png",
				Content: lang.TranslationMap{
					lang.German:  `Beschreibung 2`,
					lang.English: `Beschreibung 2`,
				},
				ContentShortLsa: lang.DefaultTranslation("short"),
				ExternalLinks:   []string{"http://google.de"},
			},
		},
		Award: []genmodel.Award{
			{
				Id:              Awa_BSS,
				Title:           lang.DefaultTranslation("Stipendium"),
				Institute:       lang.DefaultTranslation("XXX"),
				Date:            "2016",
				Content:         lang.DefaultTranslation("XXX"),
				ContentShortLsa: lang.DefaultTranslation("bester Abschluss im Fachbereich"),
				DocumentLinks:   []string{"http://google.de"},
				ExternalLinks:   []string{"http://google.de"},
			},
			{
				Id:              Awa_BSS,
				Title:           lang.DefaultTranslation("Urkunde"),
				Institute:       lang.DefaultTranslation("XXX"),
				Date:            "2017",
				Content:         lang.DefaultTranslation("XXX"),
				ContentShortLsa: lang.DefaultTranslation("short"),
				DocumentLinks:   []string{"http://google.de"},
			},
		},

		CustomPublicationLabel: lang.TranslationMap{
			lang.German:  `Publikationen`,
			lang.English: `Publications`,
		},
	}
}
