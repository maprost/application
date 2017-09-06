package genmodel

type JobAddress struct {
	Street  string
	Zip     string
	City    string
	Country string
}

type FutureExperience struct {
	JobPosition string
	StartTime   string
	Description string
	TechStack   string
}

func (e FutureExperience) Empty() bool {
	return len(e.JobPosition) == 0
}

type JobPosition struct {
	Company            string
	Address            JobAddress
	Title              string           // of the job
	MotivationText     string           //  your job specific motivation text, this text can contains tex elements
	ProfessionalSkills []SkillID        // if nothing is selected, it will use everything from profile
	SoftSkills         []SkillID        // if nothing is selected, it will use everything from profile
	MainColor          string           // please use the HTML color signature: 800000, if the field is empty: standard color will used
	FutureExperience   FutureExperience // if set, this experience will be the top of your experience
}
