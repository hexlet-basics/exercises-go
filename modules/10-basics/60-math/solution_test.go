package solution_test

import (
	"strconv"
	"testing"

	solution "exercises-go/modules/10-basics/60-math"
)

func TestMinInt(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    5,
			y:    20,
			want: 5,
		},
		{
			x:    -30,
			y:    30,
			want: -30,
		},
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    2,
			y:    0,
			want: 0,
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := solution.MinInt(tt.x, tt.y); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
