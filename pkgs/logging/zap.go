package logging

import (
	"OnlineStoreBackend/pkgs/config"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (logger *zap.Logger, err error) {
	cfg, cfgErr := config.Load([]string{"config.yaml"}, true, nil)
	if cfgErr != nil {
		panic(cfgErr)
	}

	logger, loggerErr := Configure(&cfg.Log)

	return logger, loggerErr
}

func Configure(cfg *config.Log) (logger *zap.Logger, err error) {
	if cfg == nil {
		cfg = &config.Log{}
	}

	var (
		cores   []zapcore.Core
		encoder zapcore.Encoder
		level   zapcore.Level
	)

	if cfg.File.Path != "" {
		file, err := os.OpenFile(cfg.File.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		fileWS := zapcore.Lock(file)

		level = cfg.Level
		if cfg.File.Level != nil {
			level = *cfg.File.Level
		}

		levelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= level
		})

		switch cfg.File.Encoding {
		case "json", "":
			encoderConfig := zap.NewProductionEncoderConfig()
			encoderConfig.StacktraceKey = "stacktrace"
			encoderConfig.CallerKey = "caller"
			encoder = zapcore.NewJSONEncoder(encoderConfig)
		case "console":
			encoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
		default:
			return nil, fmt.Errorf("unknown file encoding type %s", cfg.File.Encoding)
		}

		cores = append(cores, zapcore.NewCore(encoder, fileWS, levelEnabler))
	}

	if !cfg.Console.Disable {
		priorityHigh := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})

		priorityLow := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel && lvl >= level
		})

		level = cfg.Level
		if cfg.Console.Level != nil {
			level = *cfg.Console.Level
		}

		console := zapcore.Lock(os.Stdout)
		consoleErr := zapcore.Lock(os.Stderr)

		switch cfg.Console.Encoding {
		case "json":
			encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		case "console", "":
			encoderConfig := zap.NewDevelopmentEncoderConfig()
			encoderConfig.StacktraceKey = "stacktrace"
			encoderConfig.CallerKey = "caller"
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		default:
			return nil, fmt.Errorf("unknown console encoding type %s", cfg.File.Encoding)
		}

		cores = append(cores, zapcore.NewCore(encoder, consoleErr, priorityHigh))
		cores = append(cores, zapcore.NewCore(encoder, console, priorityLow))
	}

	core := zapcore.NewTee(cores...)

	stacktraceEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	logger = zap.New(core, zap.AddStacktrace(stacktraceEnabler))

	logger.Debug("configured logger", zap.String("configured-level", cfg.Level.String()))

	zap.ReplaceGlobals(logger)

	return logger, nil
}
