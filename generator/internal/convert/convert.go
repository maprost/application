package convert

import (
	"github.com/maprost/application/generator/genmodel"
)

func Convert(profile *genmodel.Profile) error {
	var err error

	// first prepare the skills
	err = setCVSkills(profile)

	return err
}
