package solution

// BEGIN

func MakeGreeting(greeting string) func(string) string {
	return func(name string) string {
		return greeting + ", " + name + "!"
	}
}

// END
