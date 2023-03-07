package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumWorker(t *testing.T) {
	a := assert.New(t)

	cat := Cat{} 
  	dog := Dog{}
  	cow := Cow{}

	a.Equal("Мяу", cat.Voice())

	a.Equal("Гав", dog.Voice())

	a.Equal("Мууу", cow.Voice())
}
