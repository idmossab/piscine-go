package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	str1 := "x = a, y = b\n"
	for _, r := range str1 {
		if r == 'a' {
			PrtPoint(points.x)
		} else if r == 'b' {
			PrtPoint(points.y)
		} else {
			z01.PrintRune(r)
		}
	}
}

func check(nb int) {
	c := '0'
	if nb == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= nb%10; i++ {
		c++
	}
	for i := -1; i >= nb%10; i-- {
		c++
	}
	if nb/10 != 0 {
		check(nb / 10)
	}
	z01.PrintRune(c)
}

func PrtPoint(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	check(n)
}
