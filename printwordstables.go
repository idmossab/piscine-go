package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _,r:=range a{
		for _,c:=range r{
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
