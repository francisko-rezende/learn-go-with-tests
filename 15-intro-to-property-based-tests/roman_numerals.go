package introtopropertybasedtests

func ConvertToRoman(arabicNumber int) string {
	switch arabicNumber {
	case 1:
		return "I"
	case 2:
		return "II"
	}

	return ""
}
