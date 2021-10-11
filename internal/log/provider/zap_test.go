package provider

import (
	"testing"
)

func TestZapDebugSuccess(t *testing.T) {
	logger := NewZap()

	logger.Debugf("test subject")
}
