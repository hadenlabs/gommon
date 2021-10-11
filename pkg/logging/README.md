# Logging

Logging.

## Installation

```sh
go get github.com/hadenlabs/gommon/pkg/logging
```

## [Usage](./logging_test.go)

```sh
import github.com/hadenlabs/gommon/pkg/logging
```

### Debugf

```go
logger := logging.New("zap")
logger.Debugf("test", "key", "value")
```

### Infof

```go
logger := logging.New("zap")
logger.Infof("test", "key", "value")
```

### Error

```go
logger := logging.New("zap")
logger.Error("Message Error")
```
