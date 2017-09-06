package lang

type TranslationMap map[Language]string

func DefaultTranslation(value string) TranslationMap {
	return TranslationMap{
		DefaultLanguage: value,
	}
}

func (t TranslationMap) String(lang Language) string {
	// try used language
	value, ok := t[lang]
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

func JoinTranslationMap(t []TranslationMap, lang Language, sep string) (result string) {
	for i, tm := range t {
		if i != 0 {
			result += sep
		}
		result += tm.String(lang)
	}
	return
}
