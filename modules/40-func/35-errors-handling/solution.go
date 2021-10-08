package solution

import (
	"errors"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

// BEGIN

var criticalErrs = []error{errBadRequest, errBadConnection}

// GetErrorMsg returns the err message if the error is critical. Otherwise it returns an empty string.
func GetErrorMsg(err error) string {
	for _, crErr := range criticalErrs {
		if errors.Is(err, crErr) {
			return crErr.Error()
		}
	}

	if errors.As(err, &nonCriticalError{}) {
		return ""
	}

	return unknownErrorMsg
}

// END
