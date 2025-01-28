package genmodel

type LeftSideActionType int

const (
	TechSkill = LeftSideActionType(iota) + 1
	Interests
	SoftSkills
	Languages
	Hobbies
	TimeAmount
	Quotation
	PersonalGoals
	MoneyAmount
	Lsa_publication
	Lsa_award
)

func (x LeftSideActionType) String() string {
	switch x {
	case TechSkill:
		return "TechSkill"
	case Interests:
		return "Interests"
	case SoftSkills:
		return "SoftSkills"
	case Languages:
		return "Languages"
	case Hobbies:
		return "Hobbies"
	case TimeAmount:
		return "TimeAmount"
	case MoneyAmount:
		return "MoneyAmount"
	case Lsa_publication:
		return "Lsa_publication"
	case Lsa_award:
		return "Lsa_award"
	default:
		return "N/A"
	}
}

type LeftSideActionTypes []LeftSideActionType

func (x LeftSideActionTypes) Index(l LeftSideActionType) (int, bool) {
	for i, a := range x {
		if a == l {
			return i, true
		}
	}
	return 0, false
}
