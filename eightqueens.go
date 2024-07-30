package piscine

import "github.com/01-edu/z01"

var (
	n     int = 8
	table []int
)

func EightQueens() {
	if n == 0 {
		for j := 0; j < len(table); j++ {
			z01.PrintRune(rune(table[j] + int('0')))
		}
		z01.PrintRune('\n')
		table = table[:len(table)-1]
		n++
		return
	}
	for i := 1; i <= 8; i++ {
		j := 0
		for j = 0; j < len(table); j++ {
			if table[j] == i || table[j] == i+n+j-8 || table[j] == 8-n-j+i {
				break
			}
		}
		if j == len(table) {
			table = append(table, i)
			n--
			EightQueens()
		}
	}
	if n < 8 {
		table = table[:len(table)-1]
		n++
	}
}
