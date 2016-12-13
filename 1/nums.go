package main

/**
First attempts at using Go...
*/
func main() {
	print(simpleSum(), "\n", smarterSum())
}

func simpleSum() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if (i % 3) == 0 {
			sum += i
			continue
		}
		if (i % 5) == 0 {
			sum += i
			continue
		}
	}

	return (sum)
}

/**
This is a more efficient solution making use of some simple Maths, made easier by using 990 instead of 999 as n
*/
func smarterSum() int {
	var sum3, sum5, sum15 int

	numTerms3 := 990 / 3
	numTerms5 := 990 / 5
	numTerms15 := 990 / 15

	sum3 = arithmeticSumToN(numTerms3, 3, 990)
	sum5 = arithmeticSumToN(numTerms5, 5, 990)
	sum15 = arithmeticSumToN(numTerms15, 15, 990)

	return sum3 + sum5 - sum15 + 993 + 995 + 996 + 999
}

func arithmeticSumToN(n, firstTerm, nthTerm int) int {
	return (n * (firstTerm + nthTerm)) / 2
}
