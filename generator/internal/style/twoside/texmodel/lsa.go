package texmodel

type RatingLsa struct {
	Name   string
	Rating int // Rating range [1,10]
}

type Language struct {
	Name  string
	Level string
}

type RangeLsa struct {
	Name            string
	Min             float64 // range [0,1]
	MinLabel        string
	MinString       string  // needed for .tex
	Max             float64 // range [0,1]
	MaxLabel        string
	MaxString       string // needed for .tex
	DeltaMaxMin     string
	DeltaMaxMinHalf string
	DeltaMaxFull    string
	SingleLabel     string
}

type RightSideAction struct {
	Type         int // 1: Skills + List, 2: List, 3: Language
	Label        string
	Ratings      []RatingLsa
	OtherRatings string
	List         string
	Languages    []Language
	Range        []RangeLsa
}
