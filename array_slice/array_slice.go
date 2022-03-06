package arrayslice

func Sum(numbers []int) int {
	result := 0

	for _, value := range numbers {
		result += value
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToGetTail ...[]int) []int {
	var result []int

	for _, numbers := range numbersToGetTail {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		}
	}

	return result
}
