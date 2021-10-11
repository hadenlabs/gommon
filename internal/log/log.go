package log

import (
	"github.com/hadenlabs/gommon/internal/errors"
	"github.com/hadenlabs/gommon/internal/log/provider"
)

// New initialize a new Log.
func NewLog(name string) TracingLogger {
	return Factory(name)
}

// Factory Log.
func Factory(name string) (prov TracingLogger) {
	switch name {
	case "zap":
		prov = provider.NewZap()
	default:
		panic(errors.Errorf(errors.ErrorParseConfig, "unsupported email provider: %s", name))
	}
	return prov
}
