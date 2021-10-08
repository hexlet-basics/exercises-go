package solution

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetErrorMsg(t *testing.T) {
	a := assert.New(t)
	a.Equal(errBadRequest.Error(), GetErrorMsg(errBadRequest))
	a.Equal(errBadConnection.Error(), GetErrorMsg(errBadConnection))
	a.Equal("", GetErrorMsg(nonCriticalError{}))
	a.Equal(unknownErrorMsg, GetErrorMsg(errors.New("unknown")))
	a.Equal(errBadRequest.Error(), GetErrorMsg(fmt.Errorf("wrap: %w", errBadRequest)))
	a.Equal(errBadConnection.Error(), GetErrorMsg(fmt.Errorf("wrap: %w", errBadConnection)))
	a.Equal("", GetErrorMsg(fmt.Errorf("wrap: %w", nonCriticalError{})))
}
