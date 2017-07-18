package santa

import (
	"github.com/maprost/application/example/max/profile"
	"github.com/maprost/application/generator/genmodel"
)

func New() genmodel.JobPosition {
	return genmodel.JobPosition{
		Title:           "Santa Clause",
		TechnicalSkills: []genmodel.SkillID{profile.TechSkillWrapping, profile.TechSkillSneaking, profile.TechSkillClimbing},
	}
}
