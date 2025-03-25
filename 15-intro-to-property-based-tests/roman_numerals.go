package introtopropertybasedtests

import "strings"

func ConvertToRoman(arabicNumber int) string {
	var result strings.Builder

	for i := arabicNumber; i > 0; i-- {
		if i == 5 {
			result.WriteString("V")
			break
		}

		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
