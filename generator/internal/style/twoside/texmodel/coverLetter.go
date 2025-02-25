package texmodel

type Address struct {
	Name     string
	Phone    string
	Email    string
	Street   string
	Zip      string
	CityOnly string
}

type CoverLetter struct {
	JobAddress     Address
	MyAddress      Address
	Subject        string
	Date           string
	TextIntro      string
	Text           string
	TextOutro      string
	Sign           string // path
	AttachmentList []string
}
