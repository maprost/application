package util

func JoinStrings(valueA string, sep string, valueB string) string {
	if valueA == "" {
		return valueB
	}
	if valueB == "" {
		return valueA
	}
	return valueA + sep + valueB
}
