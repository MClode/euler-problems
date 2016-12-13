package main

func main() {
	print(factor(600851475143))
}

func factor(num int) int {
	var f = 2

	for num > f {
		if num%f == 0 {
			num = num / f
			f = 2
		} else {
			f++
		}
	}
	return f
}
