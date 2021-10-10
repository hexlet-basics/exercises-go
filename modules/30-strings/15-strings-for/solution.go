package solution

// BEGIN

// shiftASCII shifts each byte in the string s and if there is an overflow of ASCII it takes a new code by module of 256.
func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	shifted := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		shifted[i] = byte(int(s[i])+step)
	}

	return string(shifted)
}

// END
