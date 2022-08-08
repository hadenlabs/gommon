# Logging

Logging.

## Installation

```sh
go get github.com/hadenlabs/gommon/pkg/errors
```

## [Usage](./errors_test.go)

```sh
import github.com/hadenlabs/gommon/pkg/errors
```

### Debugf

```go
logger := errors.New("zap")
logger.Debugf("test", "key", "value")
```

### Infof

```go
logger := errors.New("zap")
logger.Infof("test", "key", "value")
```

### Error

```go
logger := errors.New("zap")
logger.Error("Message Error")
```
