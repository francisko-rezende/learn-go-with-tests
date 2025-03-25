package introtopropertybasedtests

import "strings"

func ConvertToRoman(arabicNumber int) string {
	var result strings.Builder

	for arabicNumber > 0 {
		switch {
		case arabicNumber > 4:
			result.WriteString("V")
			arabicNumber -= 5
		case arabicNumber > 3:
			result.WriteString("IV")
			arabicNumber -= 4
		default:
			result.WriteString("I")
			arabicNumber--
		}
	}

	return result.String()
}
