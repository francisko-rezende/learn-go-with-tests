package introtopropertybasedtests

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabicNumber int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabicNumber >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabicNumber -= numeral.Value
		}
	}
	return result.String()
}
