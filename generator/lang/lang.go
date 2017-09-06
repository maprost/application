package lang

type Language int

const (
	English = Language(iota)
	German
)

const DefaultLanguage = English
