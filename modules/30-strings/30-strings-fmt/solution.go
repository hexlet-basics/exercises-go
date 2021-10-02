package solution

import "fmt"

// BEGIN

// generateSelfStory generates a self story substituting vars.
func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

// END
