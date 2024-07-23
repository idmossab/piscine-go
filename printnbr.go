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

	//res := ""
	table :=[20]int{}
	i:=0

	for n > 0 {
		//res += string(n % 10)
		table[i] = n%10
		n = n / 10
		i++
	}
	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune(table[j] +'0'))
	}
}
