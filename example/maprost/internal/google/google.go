package google

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
	"github.com/maprost/application/generator/style"
)

func Google() genmodel.JobPosition {
	return genmodel.JobPosition{
		MainColor: "4285F4",
		Style:     style.TwoSide,
		TwoSideStyle: genmodel.TwoSideStyle{
			//RemoveSkills: map[genmodel.LeftSideActionType][]genmodel.ID{
			//	genmodel.TechSkill:   {internal.TechSkill_Go},
			//	genmodel.Languages:   {internal.Lang_De},
			//	genmodel.MoneyAmount: {internal.Money_60for40, internal.Money_48for32},
			//},
			//Skills: map[genmodel.LeftSideActionType][]genmodel.ID{
			//	genmodel.TimeAmount: {internal.Time_1632},
			//},
			SideOneLSATypes: []genmodel.LeftSideActionType{
				genmodel.TechSkill,
				genmodel.TimeAmount,
				genmodel.MoneyAmount,
			},
			SideTwoLSATypes: []genmodel.LeftSideActionType{
				genmodel.SoftSkills,
				genmodel.Interests,
				genmodel.Languages,
				genmodel.Lsa_publication,
				genmodel.Lsa_award,
			},
			//Education: []genmodel.ID{
			//	internal.Edu_Comp,
			//},
			//RemoveEducation: []genmodel.ID{internal.Edu_Uni},
			//RemoveExperience: []genmodel.ID{internal.Exp_tutor},
			//SideOneExperienceSize: 2,
			//SideOneEducationSize: 1,
			//ShowExperienceParts:  1,
			//RemoveExperiencePart: []genmodel.ExperiencePart{genmodel.ExperiencePart_techStack},
			SideOneRSAItems: 9,
		},
		Lang:    lang.German,
		Company: "Google",
		Address: genmodel.JobAddress{
			Street:  "...",
			Zip:     "...",
			City:    "...",
			Country: "...",
		},
		Title:       "Software Defender",
		ProfileText: ``,
		FutureExperience: genmodel.FutureExperience{
			JobPosition: "Go Backend Developer",
			TechStack:   "Go, Docker",
			StartTime:   "Jan. 2018",
			Description: "",
		},
		OutputPath: genmodel.OutputPath(),
		FileName:   "",
		//TimeAmount: []genmodel.LeftSideAction{
		//	{Name: lang.TranslationMap{lang.English: "negotiable wish", lang.German: "vehandelbarer Wunsch"}, Min: 12, Max: 32},
		//	//{Name: lang.TranslationMap{lang.English: "negotiable wish 2", lang.German: "vehandelbarer Wunsch 2"}, Min: 6, Max: 32},
		//},
	}
}
