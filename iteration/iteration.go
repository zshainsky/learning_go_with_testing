package iteration

const repeatCount = 5

// Repeated function returns a string with 5 iterations of the input character
func Repeat(character string, numIterations int) string {
	var repeated string
	for i := 0; i < numIterations; i++ {
		repeated += character
	}
	return repeated
}
