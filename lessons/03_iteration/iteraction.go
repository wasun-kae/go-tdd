package iteration

func Repeat(textToRepeat string, times int) string {
	var repeatedText string

	for i := 0; i < times; i++ {
		repeatedText += textToRepeat
	}

	return repeatedText
}
