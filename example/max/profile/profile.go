package profile

import (
	"github.com/maprost/application/generator/genmodel"
	"runtime"
	"path"
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
		Name:  "Max Musertmann",
		Title: "Santa Clause",
		CV: genmodel.CV{
			Pic: path.Dir(file) + "/images/cv",
			Skills: []genmodel.Skill{
				{ID: SkillHiding, Name: "hiding", Rating: 3},
				{ID: SkillSneaking, Name: "sneaking", Rating: 5},
				{ID: SkillWrapping, Name: "wrapping", Rating: 3},
				{ID: SkillClimbing, Name: "climbing", Rating: 4},
			},
		},
	}
}
