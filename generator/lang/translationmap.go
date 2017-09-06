package lang

type TranslationMap map[Language]string

func DefaultTranslation(value string) TranslationMap {
	return TranslationMap{
		DefaultLanguage: value,
	}
}
