package solution

// BEGIN

// IsValid checks whether id and text are filled.
func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

// END
