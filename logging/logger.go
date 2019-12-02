package logging

import (
	"sync"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var baseLogger = zap.NewNop()
var mu = sync.Mutex{}

func setBaseLogger(logger *zap.Logger) {
	mu.Lock()
	defer mu.Unlock()
	baseLogger = logger
}

// NewLogger returns new zap logger for specified module name
func NewLogger(name string) *zap.Logger {
	mu.Lock()
	defer mu.Unlock()
	return baseLogger.Named(name)
}

// InitLogger initializes base Logger based on provided level and the mode
func InitLogger(level string, productionMode bool) error {
	var loggerConfig zap.Config
	if productionMode {
		loggerConfig = zapdriver.NewProductionConfig()
	} else {
		loggerConfig = zapdriver.NewDevelopmentConfig()
	}

	if level != "" {
		var zapLevel zapcore.Level
		if err := zapLevel.Set(level); err != nil {
			return err
		}
		loggerConfig.Level = zap.NewAtomicLevelAt(zapLevel)
	}

	// disable adding caller info to logs if DEBUG is not enabled
	loggerConfig.DisableCaller = !loggerConfig.Level.Enabled(zapcore.DebugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		return err
	}

	setBaseLogger(logger)
	return nil
}
