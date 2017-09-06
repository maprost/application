package lang

type Language int

const (
	English = Language(iota)
	German
)

const DefaultLanguage = English

func (l Language) String(t TranslationMap) string {
	// try used language
	value, ok := t[l]
	if ok {
		return value
	}

	// try DefaultLanguage
	value, ok = t[DefaultLanguage]
	if ok {
		return value
	}

	// try first in map
	if len(t) > 0 {
		for _, value = range t {
			return value
		}
	}

	// empty map -> empty string
	return ""
}

func (l Language) Join(t []TranslationMap, sep string) (result string) {
	for i, tm := range t {
		if i != 0 {
			result += sep
		}
		result += l.String(tm)
	}
	return
}
