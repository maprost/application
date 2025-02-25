package lang

import (
	"fmt"
	"time"
)

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

func (l Language) CityZip(city string, zip string) string {
	if l == German {
		return fmt.Sprintf("%s %s", zip, city)
	}
	return fmt.Sprintf("%s, %s", city, zip)
}

func (l Language) Date() string {
	now := time.Now().Local()
	if l == German {
		return now.Format("02.01.2006")
	}
	return now.Format("01/02/2006")
}
