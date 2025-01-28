package genmodel

type RightSideActionType int

const (
	Rsa_profile = RightSideActionType(iota) + 1
	Rsa_myMotivation
	Rsa_experience
	Rsa_education
	Rsa_publication
	Rsa_award
	Rsa_mainQuestion
)

func (x RightSideActionType) FirstSide() bool {
	switch x {
	case Rsa_myMotivation, Rsa_mainQuestion, Rsa_experience:
		return true
	default:
		return false
	}
}

func (x RightSideActionType) String() string {
	switch x {
	case Rsa_profile:
		return "Rsa_profile"
	case Rsa_myMotivation:
		return "Rsa_myMotivation"
	case Rsa_experience:
		return "Rsa_experience"
	case Rsa_education:
		return "Rsa_education"
	case Rsa_publication:
		return "Rsa_publication"
	case Rsa_award:
		return "Rsa_award"
	case Rsa_mainQuestion:
		return "Rsa_mainQuestion"
	default:
		return "N/A"
	}
}

type RightSideActionTypes []RightSideActionType

func (x RightSideActionTypes) Index(l RightSideActionType) (int, bool) {
	for i, a := range x {
		if a == l {
			return i, true
		}
	}
	return 0, false
}
