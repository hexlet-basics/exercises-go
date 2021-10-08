package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAgePopularity(t *testing.T) {
	a := assert.New(t)

	a.Equal(map[uint8]int{
		18: 1,
		44: 1,
		88: 1,
	}, PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 88},
	}.GetAgePopularity())

	a.Equal(map[uint8]int{
		18: 2,
		33: 2,
		10: 2,
	}, PersonList{
		{Age: 18},
		{Age: 33},
		{Age: 10},
		{Age: 18},
		{Age: 10},
		{Age: 33},
	}.GetAgePopularity())
}
