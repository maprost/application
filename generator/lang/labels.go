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
		//return "Zeitpräferenz"
		return "{\\small\\faClockO}\\,Zeitpräferenz"
	default:
		//return "Time preference"
		return "{\\small\\faClockO}\\,Time preference"
	}
}

func (l Language) MoneyAmount() string {
	switch l {
	case German:
		//return "{\\small\\faEuro}~ Gehaltswunsch"
		return "{\\small\\faEuro}\\,Vorstellung"
	default:
		//return "Desired Salary"
		return "{\\small\\faEuro}\\,Expectations"
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

func (l Language) PossibleUntil() string {
	switch l {
	case German:
		return "bis"
	default:
		return "until"
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
		return "\\textbf{verlinkte} Dokumente, Zeugnisse"
		//return "{\\color{linkcolor}verlinkte} Dokumente, Zeugnisse"
		//return "verlinkte Dokumente, Zeugnisse"
	default:
		return "linked documents, certificates"
	}
}

func (l Language) ExternalLinks() string {
	switch l {
	case German:
		return "externe Links"
	default:
		return "externe Links"
	}
}

func (l Language) LanguageReference() string {
	//return "test"
	switch l {
	case German:
		return "\\small\\textit{* gem. EU-Referenzrahmen} \\href{https://www.goethe.de/ins/de/de/uun/dln/ger.html}{\\faInfo}"
	default:
		return "\\small\\textit{* acc. EU framework} \\href{https://www.goethe.de/ins/de/de/uun/dln/ger.html}{\\faInfo}"
	}
}

func (l Language) AttachmentTitle() string {
	switch l {
	case German:
		return "Anhang und Hinweise"
	default:
		return "Attachment and Hints"
	}
}

func (l Language) AttachmentParaIntro() string {
	switch l {
	case German:
		return "Der Lebenslauf enthält interaktive Verlinkungen zu Dokumenten, Zeugnissen und anderen Nachweisen."
	default:
		return "The CV contains interactive links to documents, certificates and other evidence."
	}
}

func (l Language) AttachmentParaDocuments() string {
	switch l {
	case German:
		return "verlinken auf Dokumente und Zeugnisse \\\\(auf www.nat-pro.de/...)"
	default:
		return "link to documeents and certificates \\\\(on www.nat-pro.de/...)"
	}
}

func (l Language) AttachmentParaLinks() string {
	switch l {
	case German:
		return "verlinken auf andere Webseiten"
	default:
		return "link to other websites"
	}
}

func (l Language) AttachmentParaOutro() string {
	switch l {
	case German:
		return "Eine Auswahl relevanter Dokumente befindet sich im folgenden Anhang."
	default:
		return "A selection of relevant documents can be found in the following appendix."
	}
}
