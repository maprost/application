package genmodel

type JobPosition struct {
	Title  string    // of the job
	Skills []SkillID // Skills shouldn't be longer than 5
}
