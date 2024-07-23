package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}

	res := ""

	for n > 0 {
		res += string(n % 10)
		n = n / 10
	}
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(rune(res[i] + '0'))
	}
}
