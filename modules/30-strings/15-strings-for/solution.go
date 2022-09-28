package solution

// BEGIN
func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	shifted := ""

	for i := 0; i < len(s); i++ {
		shifted = shifted + string(s[i] + byte(step))
	}

	return shifted
}
// END
