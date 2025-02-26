package google

import (
	"github.com/maprost/application/example/maprost/internal"
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
	"github.com/maprost/application/generator/style"
)

func Google() genmodel.JobPosition {
	return genmodel.JobPosition{
		Attachment: []string{
			internal.Doc_CV,
			//internal.Doc_Plan,
		},
		MainColor: "4285F4",
		SideColor: "AAC9E7",
		//ShadowColor: "000000",
		//Color1:    "f73e59",
		//Color2:    "257af4",
		//Color3:    "ffbb33",
		//ScaleLineBG: "ffffff",
		Style: style.TwoSide,
		TwoSideStyle: genmodel.TwoSideStyle{
			ActivateCoverLetter: true,
			//CoverLetterOnSeparatePdf: true,
			//RemoveSkills: map[genmodel.LeftSideActionType][]genmodel.ID{
			//	genmodel.TechSkill:   {internal.TechSkill_Go},
			//	genmodel.Languages:   {internal.Lang_De},
			//	genmodel.MoneyAmount: {internal.Money_60for40, internal.Money_48for32},
			//},
			Skills: map[genmodel.LeftSideActionType][]genmodel.ID{
				//genmodel.TimeAmount: {internal.Time_HO16h, internal.Time_20full, internal.Time_1632},
				//genmodel.Languages: {internal.Lang_DeInvisible},
			},
			//SideOneLSATypes: []genmodel.LeftSideActionType{
			//	genmodel.TechSkill,
			//	genmodel.TimeAmount,
			//	genmodel.MoneyAmount,
			//},
			//SideTwoLSATypes: []genmodel.LeftSideActionType{
			//	genmodel.SoftSkills,
			//	genmodel.Lsa_publication,
			//	genmodel.Languages,
			//	genmodel.Lsa_award,
			//	genmodel.Interests,
			//},

			SideOneRSATypes: []genmodel.RightSideActionType{
				genmodel.Rsa_profile,
				genmodel.Rsa_mainQuestion,
				genmodel.Rsa_myMotivation,
				genmodel.Rsa_experience,
				genmodel.Rsa_publication,
				genmodel.Rsa_education,
				genmodel.Rsa_award,
			},
			SideOneRSAItems: 10,

			Experience: []genmodel.ID{internal.Exp_solar, internal.Exp_fitness, internal.Exp_tutor},
			//Education: []genmodel.ID{
			//	internal.Edu_Comp,
			//},
			//RemoveEducation: []genmodel.ID{internal.Edu_Uni},
			//RemoveExperience: []genmodel.ID{internal.Exp_tutor},
			//SideOneExperienceSize: 2,
			//SideOneEducationSize: 1,
			//ShowExperienceParts: 2,
			//RemoveExperiencePart: []genmodel.ExperiencePart{genmodel.ExperiencePart_techStack},
			//SideOneRSAItems: 9,
		},
		Lang:    lang.German,
		Company: "Google",
		Address: genmodel.JobAddress{
			Street:  "Straße 123",
			Zip:     "99999",
			City:    "Buxtehude",
			Country: "",
		},
		UseMig:       true,
		Title:        "Software Defender",
		ProfileLabel: "Hot'n'Cold",
		ProfileText:  ``,
		FutureExperience: genmodel.FutureExperience{
			//JobPosition: "Go Backend Developer",
			JobPositionFirstLine: `Go Backend Developer`,
			JobPosition:          `für ganz, ganz langer Titel`,
			//TechStack:   "Go, Docker",
			StartTime: "01/2018",
			// TODO multilanguage
			EndTime: "Vertragsende",
			//EndTime:     "12/2018",
			Description: "an diesen Aufgaben hätte ich Spaß und an diesen Aufgaben auch und an diesen sowieso an diesen Aufgaben hätte ich Spaß und an diesen Aufgaben auch und an diesen sowieso an diesen Aufgaben hätte ich Spaß und an diesen Aufgaben auch und an diesen sowieso ",
		},
		FutureColorChange: true,
		OutputPath:        genmodel.OutputPath(),
		FileName:          "",
		TechSkillsLabel:   "Fachkompetenz",
		SoftSkillsLabel:   "Sozialkompetenz",
		MainQuestionLabel: "MainQuestion Label",
		MainQuestionText:  "MainQuestion Text",

		CoverLetterSubject:  "Betrifft das Jobangebot XY",
		CoverLetterTxtIntro: `Liebes XY-Team`,
		CoverLetterTxtOutro: `Mit bestem Dank und besten Grüßen`,
		CoverLetterTxt: `hiermit bewerbe ich mich...` +
			`` +
			`\\` +
			``,
		CoverLetterAttachment: []string{"Blob", "Drop"},
	}
}
