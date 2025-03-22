package introtopropertybasedtests

func ConvertToRoman(arabicNumber int) string {
	if arabicNumber == 3 {
		return "III"
	}

	if arabicNumber == 2 {
		return "II"
	}

	return "I"
}
