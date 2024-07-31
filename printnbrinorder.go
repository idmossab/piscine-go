package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		n *= -1
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	res := ""
	for n > 0 {
		res += string(n % 10)
		n /= 10
	}
	for i := 0; i < len(res); i++ {
		z01.PrintRune(rune(res[i] + '0'))
	}
}
