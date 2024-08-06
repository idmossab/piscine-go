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
	str1 := "x = "
	for _, r := range str1 {
		z01.PrintRune(r)
	}
	PrtPoint(points.x)
	str2 := ", y = "
	for _, r1 := range str2 {
		z01.PrintRune(r1)
	}
	PrtPoint(points.y)
	z01.PrintRune('\n')
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
