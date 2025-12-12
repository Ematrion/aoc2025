package utils


func AbsInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func ModularAtithemtic(a, m int) (b, k int) {
	// positiv remainder even when a is negativ
	b = ((a % m) + m) % m
	k = (a-b) / m
	return b, k
}