package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func logForTest() (TracingLogger, func()) {
	log := NewLog("zap")

	return log, func() {}
}

func TestNewSuccess(t *testing.T) {
	log, tearDown := logForTest()
	defer tearDown()
	assert.Implements(t, (*TracingLogger)(nil), log)
}
