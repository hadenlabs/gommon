package logging

import (
	"github.com/hadenlabs/gommon/internal/log"
)

func New(name string) log.TracingLogger {
	return log.Factory(name)
}
