package piscine

func AtoiBase(s string, base string) int {
	n := 0
	if isValidate1(base) == false {
		return 0
	}
	// convert to base 10
	for _, r := range s {
		for i := range base {
			if r == rune(base[i]) {
				n *= len(base)
				n += i
			}
		}
	}
	return n
}

func isValidate1(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return false
		}
		for j := 0; j < len(base); j++ {
			if j != i && base[j] == base[i] {
				return false
			}
		}
	}
	return true
}
