package google

import (
	"github.com/maprost/application/example/maprost/internal"
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/lang"
	"github.com/maprost/application/generator/style"
)

func Google() genmodel.JobPosition {
	return genmodel.JobPosition{
		MainColor: "4285F4",
		Style:     style.TwoSide,
		TwoSideStyle: genmodel.TwoSideStyle{
			RemoveSkills: map[genmodel.LeftSideActionType][]genmodel.ID{
				genmodel.TechSkill: {internal.TechSkill_Go},
				genmodel.Languages: {internal.Lang_De},
			},
			SideTwoLeftSideActionTypes: []genmodel.LeftSideActionType{
				genmodel.SoftSkills,
			},
			//Education: []genmodel.ID{
			//	internal.Edu_Comp,
			//},
			RemoveEducation: []genmodel.ID{internal.Edu_Uni},
			//RemoveExperience: []genmodel.ID{internal.Exp_tutor},
			//SideOneExperienceSize: 2,
			SideOneEducationSize: 1,
			ShowExperienceParts:  1,
			RemoveExperiencePart: []genmodel.ExperiencePart{genmodel.ExperiencePart_techStack},
		},
		Lang:    lang.German,
		Company: "Google",
		Address: genmodel.JobAddress{
			Street:  "...",
			Zip:     "...",
			City:    "...",
			Country: "...",
		},
		Title:          "Software Defender",
		MotivationText: ``,
		FutureExperience: genmodel.FutureExperience{
			JobPosition: "Go Backend Developer",
			TechStack:   "Go, Docker",
			StartTime:   "Jan. 2018",
			Description: "",
		},
		OutputPath: genmodel.OutputPath(),
		FileName:   "",
	}
}
