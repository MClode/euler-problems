package main

func main() {
	print(fibonacci(0, 1, 0), "\n")
}

func fibonacci(x, y, total int) int {
	if (x+y)%2 == 0 {
		total += x + y
	}

	if x+y+y > 4000000 {
		return total
	}

	return fibonacci(y, x+y, total)
}
