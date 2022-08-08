package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"

	coreErrors "github.com/hadenlabs/gommon/internal/errors"

	"fmt"
)

func TestFormatSuccess(t *testing.T) {
	err := New(coreErrors.ErrorNotImplemented, "Internal error")
	assert.Equal(t, "Internal error", fmt.Sprintf("%v", err))
	assert.Contains(t, fmt.Sprintf("%+v", err), "Internal error", err) // with stacktrace
}
