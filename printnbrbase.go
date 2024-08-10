package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if isValidate(base) == false {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return

	}
	if nbr < 0 {
		z01.PrintRune('-')
		ConvertNbr(nbr/(-len(base)), base)
		z01.PrintRune(rune(base[-nbr%len(base)]))
	} else {
		ConvertNbr(nbr, base)
	}
}

func ConvertNbr(nbr int, base string) {
	// convert to another base
	converted := ""
	if nbr == 0 {
		z01.PrintRune('0')
	}
	for nbr > 0 {
		converted = string(base[nbr%len(base)]) + converted
		nbr /= len(base)
	}
	for _, s := range converted {
		z01.PrintRune(s)
	}
}

func isValidate(base string) bool {
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
