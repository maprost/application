package santa

import (
	"github.com/maprost/application/example/max/profile"
	"github.com/maprost/application/generator/genmodel"
)

func New() genmodel.Company {
	return genmodel.Company{
		Skills: []genmodel.SkillID{profile.SkillWrapping, profile.SkillSneaking, profile.SkillClimbing},
	}
}
