package main

func main() {
	x := sumOfSquares(1, 0)
	y := squareOfSum(1, 0)

	print(y - x)
}

func sumOfSquares(num, sum int) int {
	if num > 100 {
		return sum
	}
	square := num * num
	return sumOfSquares(num+1, square+sum)
}

func squareOfSum(num, sum int) int {
	if num > 100 {
		return sum * sum
	}
	return squareOfSum(num+1, sum+num)
}
