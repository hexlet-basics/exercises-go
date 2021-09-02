package solution

// BEGIN

// Valid checks whether id and text are filled.
func Valid(id int, text string) bool {
	return id > 0 && text != ""
}

// END
