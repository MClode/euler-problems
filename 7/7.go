package main

import (
	"math"
	"errors"
	"fmt"
)

func main()  {
	np := 10001
	total, err := upperBound(float64(np))

	if (err != nil) {
		fmt.Print(err)
		return
	}

	nums := make([]bool, total)

	next := 2

	MainLoop:
	for next < total {

		loop(total, next, nums)

		for x:=next+1; x < total; x++ {
			if nums[x] != true {
				next = x
				break
			}

			if (next > total/2) {
				break MainLoop
			}
		}
	}

	count := 0

	for x:=1; x < total; x++ {
		if (nums[x] == false) {
			if (count == np) {
				fmt.Print(count, " prime = ", x, "\n")
				return
			}
			count++
		}
	}
}

func loop(total, next int, nums []bool) {
	start := next * 2
	gap := next

	for x:=start; x < total; x = x + gap {
		nums[x] = true
	}
}

func upperBound(num float64) (int, error) {
	if (num < 6) { return -1, errors.New("This method should only be used for n >= 6") }

	return int(num * math.Log2(num) + num * math.Log2(math.Log2(num))) + 1, nil
}