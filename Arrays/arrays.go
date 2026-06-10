package main

func Sum(numbers []int) int{
	summation := 0

	for _, number := range numbers {
		summation += number
	}

	return summation
}
