package iteration

func Repeat(character string, timesToRepeat int) string {
	var repeated string
	for i := 0; i < timesToRepeat; i++ {
		repeated += character
	}

	return repeated
}
