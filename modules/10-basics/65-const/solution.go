package solution

// BEGIN

// error messages
const (
	OkMsg        = "OK"
	CancelledMsg = "CANCELLED"
)

// error codes
const (
	OkCode = iota
	CancelledCode
	UnknownCode
)

// ErrorMessageToCode returns a gRPC error code depending on error message.
func ErrorMessageToCode(msg string) int {
	switch msg {
	case OkMsg:
		return OkCode
	case CancelledMsg:
		return CancelledCode
	}

	return UnknownCode
}

// END
