package main

/**
 * This is an intentionally brute force method (with a big O that does not bare thinking about), as I play around with
 * concurrency in Go
 */
func main() {
	for x := 2520; x <= 1000000000; x += 20 {
		for y := 3; y <= 20; y++ {
			if x%y != 0 {
				print(x, " failed on ", y, "\n")
				break
			}

			if y == 20 {
				print(x)
				return
			}
		}
	}
}
