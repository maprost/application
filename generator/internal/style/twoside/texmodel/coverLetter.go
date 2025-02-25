package texmodel

type Address struct {
	Name   string
	Phone  string
	Email  string
	Street string
	Zip    string
}

type CoverLetter struct {
	JobAddress     Address
	MyAddress      Address
	Subject        string
	Date           string
	Text           string
	AttachmentList []string
}
