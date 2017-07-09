package convert

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/texmodel"
)

func Convert(profile *genmodel.Profile) (texmodel.Index, error) {
	var err error

	// first prepare the skills
	err = setCVSkills(profile)

	return texmodel.Index{}, err
}
