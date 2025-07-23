package solution

import "fmt"

// BEGIN

func GetFileExtension(filename string) (string, error) {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i+1:], nil
		}
	}
	return "", fmt.Errorf("файл %q не имеет расширения", filename)
}

// END
