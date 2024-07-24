package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n *=-1
	}
	if n == 0 {
		z01.PrintRune('0')
	}

	res := ""
	for n > 0 {
		res += string(n % 10)
		n = n / 10
	}
	for j := len(res) - 1; j >= 0; j-- {
		z01.PrintRune(rune(res[j] + '0'))
	}
}
