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
)

func (x LeftSideActionType) FirstSide() bool {
	switch x {
	case TechSkill, Interests, SoftSkills, Languages, Hobbies:
		return true
	default:
		return false
	}
}

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
