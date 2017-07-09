package profile

import (
	"github.com/maprost/application/generator/genmodel"
	"path"
	"runtime"
)

const (
	SkillHiding = genmodel.SkillID(iota)
	SkillSneaking
	SkillWrapping
	SkillClimbing
)

func New() genmodel.Profile {
	_, file, _, _ := runtime.Caller(1)

	return genmodel.Profile{
		Name:  "Mad Max",
		Title: "Bringer of Joy",
		CV: genmodel.CV{
			Pic: path.Dir(file) + "/image/cv",
			Skills: []genmodel.Skill{
				{ID: SkillHiding, Name: "hiding", Rating: 3},
				{ID: SkillSneaking, Name: "sneaking", Rating: 5},
				{ID: SkillWrapping, Name: "wrapping", Rating: 3},
				{ID: SkillClimbing, Name: "climbing", Rating: 4},
			},
		},
	}
}
