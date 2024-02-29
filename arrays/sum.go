package main

//below is a way to pass the test

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

//However, range lets you iterate over an array. On each iteration, range returns two values - the index and the value. We are choosing to ignore the index value by using _
// i = number
func Sum(numbers [5]int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
