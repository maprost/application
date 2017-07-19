package genmodel

type JobPosition struct {
	Company            string
	Address            Address
	Title              string    // of the job
	LetterText         string    // this text can contains tex elements
	ProfessionalSkills []SkillID // if nothing is selected, it will use everything from profile
	SoftSkills         []SkillID // if nothing is selected, it will use everything from profile
	MainColor          string    // please use the HTML color signature: 800000, if the field is empty: standard color will used
}
