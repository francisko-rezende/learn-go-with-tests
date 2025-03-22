package introtopropertybasedtests

import "strings"

func ConvertToRoman(arabicNumber int) string {
	var result strings.Builder

	for i := 0; i < arabicNumber; i++ {
		result.WriteString("I")
	}

	return result.String()
}
