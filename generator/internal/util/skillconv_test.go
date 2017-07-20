package util_test

import (
	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/assertion"
	"testing"
)

const (
	skillOneID = genmodel.SkillID(iota)
	skillTwoID
	skillThreeID
	skillFourID
)

var (
	skillOne   = genmodel.Skill{Name: "one", Rating: 1}
	skillTwo   = genmodel.Skill{Name: "two", Rating: 2}
	skillThree = genmodel.Skill{Name: "three", Rating: 3}
	skillFour  = genmodel.Skill{Name: "four", Rating: 4}
)

func TestCalculateProfessionalSkills_emptyInput(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile:     genmodel.Profile{},
		JobPosition: genmodel.JobPosition{},
	}

	skills, err := util.CalculateProfessionalSkills(&application)
	assert.Nil(err)
	assert.Len(skills, 0)
}

func TestCalculateProfessionalSkills_normalInput(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			ProfessionalSkills: []genmodel.SkillID{skillOneID, skillTwoID, skillThreeID, skillFourID},
		},
	}

	skills, err := util.CalculateProfessionalSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.Skill{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateProfessionalSkills_normalInput_noFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{},
	}

	skills, err := util.CalculateProfessionalSkills(&application)
	assert.Nil(err)

	// ordering is up to the map
	assert.Similar(skills, []genmodel.Skill{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateProfessionalSkills_normalInput_someFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			ProfessionalSkills: []genmodel.SkillID{skillTwoID, skillThreeID},
		},
	}

	skills, err := util.CalculateProfessionalSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.Skill{
		skillTwo, skillThree,
	})
}

func TestCalculateProfessionalSkills_normalInput_wrongFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
			},
		},
		JobPosition: genmodel.JobPosition{
			ProfessionalSkills: []genmodel.SkillID{skillOneID, skillFourID},
		},
	}

	_, err := util.CalculateProfessionalSkills(&application)
	assert.NotNil(err)
}

func TestCalculateSoftSkills_emptyInput(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile:     genmodel.Profile{},
		JobPosition: genmodel.JobPosition{},
	}

	skills, err := util.CalculateSoftSkills(&application)
	assert.Nil(err)
	assert.Len(skills, 0)
}

func TestCalculateSoftSkills_normalInput(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			SoftSkills: []genmodel.SkillID{skillOneID, skillTwoID, skillThreeID, skillFourID},
		},
	}

	skills, err := util.CalculateSoftSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.Skill{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateSoftSkills_normalInput_noFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{},
	}

	skills, err := util.CalculateSoftSkills(&application)
	assert.Nil(err)

	// ordering is up to the map
	assert.Similar(skills, []genmodel.Skill{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateSoftSkills_normalInput_someFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			SoftSkills: []genmodel.SkillID{skillTwoID, skillThreeID},
		},
	}

	skills, err := util.CalculateSoftSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.Skill{
		skillTwo, skillThree,
	})
}

func TestCalculateSoftSkills_normalInput_wrongFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.SkillID]genmodel.Skill{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
			},
		},
		JobPosition: genmodel.JobPosition{
			SoftSkills: []genmodel.SkillID{skillOneID, skillFourID},
		},
	}

	_, err := util.CalculateSoftSkills(&application)
	assert.NotNil(err)
}
