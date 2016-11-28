package main

import "strconv"

/**
 * Brute force & ugly solution!
 */
func main() {
	Palindrome := 0
	for x := 900; x <= 999; x++ {
		for y := 900; y <= 999; y++ {
			sum := x * y
			if (isPalindrome(strconv.Itoa(sum)) && (sum) > Palindrome) {
				Palindrome = sum
			}
		}
	}

	print(Palindrome)
}

func isPalindrome(s string) bool {

	if (s[len(s)-1] == s[0]) {
		s = s[1:len(s)-1]
	} else {
		return false
	}

	if (len(s) < 2) {
		return true
	}

	return isPalindrome(s)
}
