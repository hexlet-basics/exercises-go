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
