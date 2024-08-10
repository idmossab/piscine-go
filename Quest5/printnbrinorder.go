package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	res := ""
	for n > 0 {
		res += string(rune(n%10 + '0'))
		n /= 10
	}
	runes := []rune(res)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] > runes[i+1] {
			runes[i], runes[i+1] = runes[i+1], runes[i]
			i = -1
		}
	}
	for i := 0; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
}
