package convert_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/application/generator/genmodel"
	"github.com/maprost/application/generator/internal/convert"
)

const (
	SkillSleeping = genmodel.SkillID(iota)
	SkillDrinkingCoffee
	SkillWatchingYoutube
	SkillChat
	SkillLookingOutOfTheWindow
	SkillGoingInTheBathroom
)

func TestConvert(t *testing.T) {
	assert := assertion.New(t)

	profile := genmodel.Profile{
		CV: genmodel.CV{
			Skills: []genmodel.Skill{
				{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
				{ID: SkillDrinkingCoffee, Name: "Drinking Coffee", Rating: 3},
				{ID: SkillWatchingYoutube, Name: "Watching Youtube", Rating: 5},
				{ID: SkillChat, Name: "Chat", Rating: 2},
				{ID: SkillLookingOutOfTheWindow, Name: "Looking out of the window", Rating: 5},
				{ID: SkillGoingInTheBathroom, Name: "Going to the bathroom", Rating: 2},
			},
		},
		Company: genmodel.Company{
			Skills: []genmodel.SkillID{SkillSleeping},
		},
	}

	err := convert.Convert(&profile)
	assert.Nil(err)
	assert.Equal(profile.CV.Skills, []genmodel.Skill{
		{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
	})
}

func TestConvert_toMuchSkillsInCompany_cutToFive(t *testing.T) {
	assert := assertion.New(t)

	profile := genmodel.Profile{
		CV: genmodel.CV{
			Skills: []genmodel.Skill{
				{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
				{ID: SkillDrinkingCoffee, Name: "Drinking Coffee", Rating: 3},
				{ID: SkillWatchingYoutube, Name: "Watching Youtube", Rating: 5},
				{ID: SkillChat, Name: "Chat", Rating: 2},
				{ID: SkillLookingOutOfTheWindow, Name: "Looking out of the window", Rating: 5},
				{ID: SkillGoingInTheBathroom, Name: "Going to the bathroom", Rating: 2},
			},
		},
		Company: genmodel.Company{
			Skills: []genmodel.SkillID{
				SkillSleeping,
				SkillDrinkingCoffee,
				SkillWatchingYoutube,
				SkillChat,
				SkillLookingOutOfTheWindow,
				SkillGoingInTheBathroom,
			},
		},
	}

	err := convert.Convert(&profile)
	assert.Nil(err)
	assert.Equal(profile.CV.Skills, []genmodel.Skill{
		{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
		{ID: SkillDrinkingCoffee, Name: "Drinking Coffee", Rating: 3},
		{ID: SkillWatchingYoutube, Name: "Watching Youtube", Rating: 5},
		{ID: SkillChat, Name: "Chat", Rating: 2},
		{ID: SkillLookingOutOfTheWindow, Name: "Looking out of the window", Rating: 5},
	})
}

func TestConvert_lessThanFiveProfileSkills(t *testing.T) {
	assert := assertion.New(t)

	profile := genmodel.Profile{
		CV: genmodel.CV{
			Skills: []genmodel.Skill{
				{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
				{ID: SkillDrinkingCoffee, Name: "Drinking Coffee", Rating: 3},
			},
		},
		Company: genmodel.Company{
			Skills: []genmodel.SkillID{
				SkillSleeping,
			},
		},
	}

	err := convert.Convert(&profile)
	assert.Nil(err)
	assert.Equal(profile.CV.Skills, []genmodel.Skill{
		{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
	})
}

func TestConvert_noCompanySkillSetup_chooseFirstFiveProfileSkills(t *testing.T) {
	assert := assertion.New(t)

	profile := genmodel.Profile{
		CV: genmodel.CV{
			Skills: []genmodel.Skill{
				{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
				{ID: SkillDrinkingCoffee, Name: "Drinking Coffee", Rating: 3},
				{ID: SkillWatchingYoutube, Name: "Watching Youtube", Rating: 5},
				{ID: SkillChat, Name: "Chat", Rating: 2},
				{ID: SkillLookingOutOfTheWindow, Name: "Looking out of the window", Rating: 5},
				{ID: SkillGoingInTheBathroom, Name: "Going to the bathroom", Rating: 2},
			},
		},
		Company: genmodel.Company{},
	}

	err := convert.Convert(&profile)
	assert.Nil(err)
	assert.Equal(profile.CV.Skills, []genmodel.Skill{
		{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
		{ID: SkillDrinkingCoffee, Name: "Drinking Coffee", Rating: 3},
		{ID: SkillWatchingYoutube, Name: "Watching Youtube", Rating: 5},
		{ID: SkillChat, Name: "Chat", Rating: 2},
		{ID: SkillLookingOutOfTheWindow, Name: "Looking out of the window", Rating: 5},
	})
}

func TestConvert_unknownCompanySkill_error(t *testing.T) {
	assert := assertion.New(t)

	profile := genmodel.Profile{
		CV: genmodel.CV{
			Skills: []genmodel.Skill{
				{ID: SkillSleeping, Name: "Sleeping", Rating: 4},
			},
		},
		Company: genmodel.Company{
			Skills: []genmodel.SkillID{
				SkillSleeping,
				SkillDrinkingCoffee,
			},
		},
	}

	err := convert.Convert(&profile)
	assert.NotNil(err)
}
