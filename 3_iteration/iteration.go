package iteration

func RepeatChar(character string, repetitions int) string {
	var repeated string
	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return repeated
}
