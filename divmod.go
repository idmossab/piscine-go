package piscine

func DivMod(a int, b int, div *int, mod *int) {
	d := a/b
	m := a%b
	*div = d
	*mod = m
}