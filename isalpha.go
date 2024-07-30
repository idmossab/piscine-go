package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if check(runes[i]) == false {
			return false
		}
	}
	return true
}

func check(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}
