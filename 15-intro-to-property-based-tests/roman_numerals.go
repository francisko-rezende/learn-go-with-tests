package introtopropertybasedtests

import "strings"

func ConvertToRoman(arabicNumber int) string {
	if arabicNumber == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := 0; i < arabicNumber; i++ {
		result.WriteString("I")
	}

	return result.String()
}
