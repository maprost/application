package bunny

import (
	"github.com/maprost/application/example/max/profile"
	"github.com/maprost/application/generator/genmodel"
)

func New() genmodel.JobPosition {
	return genmodel.JobPosition{
		Skills: []genmodel.SkillID{profile.SkillHiding, profile.SkillSneaking},
	}
}
