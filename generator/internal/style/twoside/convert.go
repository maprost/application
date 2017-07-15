package twoside

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/style/twoside/texmodel"
)

func convert(profile *genmodel.Profile) (texmodel.Index, error) {
	var err error

	// first prepare the skills
	err = setCVSkills(profile)

	return texmodel.Index{}, err
}
