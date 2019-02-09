package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

//...[]int is variadic syntax. Can take variable num of args
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNums := len(numbersToSum)

	sums := make([]int, lengthOfNums)

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}
	return sums
}
