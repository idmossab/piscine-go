package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}
	res := 1
	for nb > 0 {
		res *= nb
		nb--
	}
	return res
}
