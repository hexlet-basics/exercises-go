package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSelfStory(t *testing.T) {
	a := assert.New(t)
	a.Equal("Hello! My name is Vlad. I'm 25 y.o. And I also have $10.00 in my wallet right now.", generateSelfStory("Vlad", 25, 10.00000025))
}
