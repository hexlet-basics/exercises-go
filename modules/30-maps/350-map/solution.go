package solution

import (
	"errors"
	"fmt"
)

// BEGIN

func GetGrade(scores map[string]int, name string) (string, error) {
	if score, exists := scores[name]; exists {
		return fmt.Sprintf("%s has %d points", name, score), nil
	}
	return "", errors.New("student not found")
}

// END
