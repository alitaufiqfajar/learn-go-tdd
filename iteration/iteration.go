package iteration

const repeatCount = 5

func Repeat(character string, count int) string {
	result := ""
	cycles := repeatCount

	if count > 0 {
		cycles = count
	}

	for i := 0; i < cycles; i++ {
		result += character
	}

	return result
}
