package lang

func (l Language) Languages() string {
	switch l {
	case German:
		return "Sprachen"
	default:
		return "Languages"
	}
}

func (l Language) Interests() string {
	switch l {
	case German:
		return "Interessen"
	default:
		return "Interests"
	}
}

func (l Language) SoftSkills() string {
	switch l {
	case German:
		return "Soft\\,Skills"
	default:
		return "Soft\\,Skills"
	}
}

func (l Language) TechSkills() string {
	switch l {
	case German:
		return "Tech.\\,Skills"
	default:
		return "Tech.\\,Skills"
	}
}

func (l Language) Hobbies() string {
	switch l {
	case German:
		return "Freizeit"
	default:
		return "Hobbies"
	}
}

func (l Language) Profile() string {
	switch l {
	case German:
		return "Profile"
	default:
		return "Profile"
	}
}

func (l Language) Education() string {
	switch l {
	case German:
		return "Bildung"
	default:
		return "Education"
	}
}

func (l Language) Experience() string {
	switch l {
	case German:
		return "Erfahrung"
	default:
		return "Experience"
	}
}

func (l Language) Since() string {
	switch l {
	case German:
		return "seit"
	default:
		return "since"
	}
}

func (l Language) PossibleAt() string {
	switch l {
	case German:
		return "verfügbar ab"
	default:
		return "possible at"
	}
}
