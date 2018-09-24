package diffsquares

//SumOfSquares returns sum of squares of n natural numbers
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

//SquareOfSum returns the sq. of sum of n natural numbers
func SquareOfSum(n int) int {
	var sum int
	sum += (n * (n + 1)) / 2
	return sum * sum
}

//Difference returns diff between sum f sq and sq of sums of n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
