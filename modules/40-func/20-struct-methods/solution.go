package solution

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// BEGIN
func (c *Counter) Inc(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Value = Max(c.Value+delta, 0)
}

func (c *Counter) Dec(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Inc(-delta)
}

// END
