package piscine

func IterativePower(nb int, power int) int {
	if power == 0 || nb == 1 {
		return 1
	}
	res := 1
	for power > 0 {
		res *= nb
		power--
	}
	return res
}
