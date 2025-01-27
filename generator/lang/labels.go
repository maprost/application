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

func (l Language) TimeAmount() string {
	switch l {
	case German:
		return "Zeitpräferenz"
	default:
		return "Time preference"
	}
}

func (l Language) MoneyAmount() string {
	switch l {
	case German:
		return "Gehaltswunsch"
	default:
		return "Desired Salary"
	}
}

func (l Language) Profile() string {
	switch l {
	case German:
		return "Profil"
	default:
		return "Profile"
	}
}

func (l Language) Motivation() string {
	switch l {
	case German:
		return "Motivation"
	default:
		return "Motivation"
	}
}

func (l Language) MainQuestion() string {
	switch l {
	case German:
		return "Aber warum..."
	default:
		return "But why..."
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

func (l Language) Publication() string {
	switch l {
	case German:
		return "Veröffentlichungen"
	default:
		return "Publications"
	}
}

func (l Language) Award() string {
	switch l {
	case German:
		return "Auszeichnungen"
	default:
		return "Awards"
	}
}

func (l Language) FinalGrade() string {
	switch l {
	case German:
		return "Abschlussnote"
	default:
		return "Final grade"
	}
}

func (l Language) QuitReason() string {
	switch l {
	case German:
		return ""
	default:
		return ""
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

func (l Language) SomeProjects() string {
	switch l {
	case German:
		return "ausgewählte Projekte"
	default:
		return "some projects"
	}
}

func (l Language) SomeRoles() string {
	switch l {
	case German:
		return "ausgewählte Rollen"
	default:
		return "some roles"
	}
}

func (l Language) UsedTechs() string {
	switch l {
	case German:
		return "genutzte Technologien"
	default:
		return "used technologies"
	}
}

func (l Language) DocumentLinks() string {
	switch l {
	case German:
		return "Dokumente, Zeugnisse"
	default:
		return "Documents, certificates"
	}
}
