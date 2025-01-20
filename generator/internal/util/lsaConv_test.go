package util_test

import (
	"testing"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/util"
	"github.com/maprost/assertion"
)

const (
	skillOneID = genmodel.ID(iota)
	skillTwoID
	skillThreeID
	skillFourID
)

var (
	skillOne   = genmodel.LeftSideAction{Name: "one", Rating: 1}
	skillTwo   = genmodel.LeftSideAction{Name: "two", Rating: 2}
	skillThree = genmodel.LeftSideAction{Name: "three", Rating: 3}
	skillFour  = genmodel.LeftSideAction{Name: "four", Rating: 4}
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
			ProfessionalSkills: map[genmodel.ID]genmodel.LeftSideAction{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			ProfessionalSkills: []genmodel.ID{skillOneID, skillTwoID, skillThreeID, skillFourID},
		},
	}

	skills, err := util.CalculateProfessionalSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.LeftSideAction{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateProfessionalSkills_normalInput_noFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.ID]genmodel.LeftSideAction{
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
	assert.Similar(skills, []genmodel.LeftSideAction{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateProfessionalSkills_normalInput_someFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.ID]genmodel.LeftSideAction{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			ProfessionalSkills: []genmodel.ID{skillTwoID, skillThreeID},
		},
	}

	skills, err := util.CalculateProfessionalSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.LeftSideAction{
		skillTwo, skillThree,
	})
}

func TestCalculateProfessionalSkills_normalInput_wrongFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			ProfessionalSkills: map[genmodel.ID]genmodel.LeftSideAction{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
			},
		},
		JobPosition: genmodel.JobPosition{
			ProfessionalSkills: []genmodel.ID{skillOneID, skillFourID},
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
			SoftSkills: map[genmodel.ID]genmodel.LeftSideAction{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			SoftSkills: []genmodel.ID{skillOneID, skillTwoID, skillThreeID, skillFourID},
		},
	}

	skills, err := util.CalculateSoftSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.LeftSideAction{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateSoftSkills_normalInput_noFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.ID]genmodel.LeftSideAction{
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
	assert.Similar(skills, []genmodel.LeftSideAction{
		skillOne, skillTwo, skillThree, skillFour,
	})
}

func TestCalculateSoftSkills_normalInput_someFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.ID]genmodel.LeftSideAction{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
				skillFourID:  skillFour,
			},
		},
		JobPosition: genmodel.JobPosition{
			SoftSkills: []genmodel.ID{skillTwoID, skillThreeID},
		},
	}

	skills, err := util.CalculateSoftSkills(&application)
	assert.Nil(err)
	assert.Equal(skills, []genmodel.LeftSideAction{
		skillTwo, skillThree,
	})
}

func TestCalculateSoftSkills_normalInput_wrongFilter(t *testing.T) {
	assert := assertion.New(t)

	application := genmodel.Application{
		Profile: genmodel.Profile{
			SoftSkills: map[genmodel.ID]genmodel.LeftSideAction{
				skillOneID:   skillOne,
				skillTwoID:   skillTwo,
				skillThreeID: skillThree,
			},
		},
		JobPosition: genmodel.JobPosition{
			SoftSkills: []genmodel.ID{skillOneID, skillFourID},
		},
	}

	_, err := util.CalculateSoftSkills(&application)
	assert.NotNil(err)
}
