package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"

	customLog "github.com/hadenlabs/gommon/internal/log"
)

func logForTest() (customLog.TracingLogger, func()) {
	logger := New("zap")

	return logger, func() {}
}

func TestNewSuccess(t *testing.T) {
	logger, tearDown := logForTest()
	defer tearDown()
	assert.Implements(t, (*customLog.TracingLogger)(nil), logger)
}

func TestDebugfSuccess(t *testing.T) {
	log, tearDown := logForTest()
	defer tearDown()
	assert.NotPanicsf(t, func() {
		log.Debugf("test", "key", "value")
	}, "implement function debugf")
}

func TestInfofSuccess(t *testing.T) {
	log, tearDown := logForTest()
	defer tearDown()
	assert.NotPanicsf(t, func() {
		log.Infof("test", "key", "value")
	}, "implement function debugf")
}

func TestErrorSuccess(t *testing.T) {
	log, tearDown := logForTest()
	defer tearDown()
	assert.NotPanicsf(t, func() {
		log.Error("test")
	}, "implement function debugf")
}
