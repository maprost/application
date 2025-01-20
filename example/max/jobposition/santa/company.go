package santa

import (
	"github.com/maprost/application/generator/genmodel"
)

func New() genmodel.JobPosition {
	return genmodel.JobPosition{
		Title: "Santa Clause",
		//ProfessionalSkills: []genmodel.ID{profile.TechSkillWrapping, profile.TechSkillSneaking, profile.TechSkillClimbing},
	}
}
