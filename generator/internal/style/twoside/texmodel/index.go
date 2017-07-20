package texmodel

type Index struct {
	FirstPage    FirstPage   // only used in long version
	CoverLetter  CoverLetter // only used in long version
	CV           CV
	MainColor    string
	ShortVersion bool
}
