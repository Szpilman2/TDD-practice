package main

func Sum(numbers []int) int{
	summation := 0

	for _, number := range numbers {
		summation += number
	}

	return summation
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}