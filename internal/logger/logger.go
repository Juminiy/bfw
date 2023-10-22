package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

var (
	LogLevelMapping = map[string]zapcore.Level{
		"DEBUG":  zapcore.DebugLevel,
		"INFO":   zapcore.InfoLevel,
		"WARN":   zapcore.WarnLevel,
		"ERROR":  zapcore.ErrorLevel,
		"DPANIC": zapcore.DPanicLevel,
		"PANIC":  zapcore.PanicLevel,
		"FATAL":  zapcore.FatalLevel,
	}
	_sugar       *zap.SugaredLogger
	_logger      *zap.Logger
	_initialized = false
)

func HiddenCallerEncoder(_ zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("hidden")
}

func InitLogger(logPath string, logLevel string) {
	var err error
	if strings.Contains(logPath, "/") {
		if ok, _ := HasDir(logPath[:strings.LastIndex(logPath, "/")]); !ok {
			err = os.MkdirAll(logPath[:strings.LastIndex(logPath, "/")], os.ModePerm)
		}
		if err != nil {
			log.Fatalf("failed to create logger path: %v", err)
			os.Exit(-2)
		}
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		logPath,
	}
	cfg.ErrorOutputPaths = []string{
		logPath,
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeCaller = HiddenCallerEncoder
	cfg.Level = zap.NewAtomicLevelAt(LogLevelMapping[logLevel])

	_logger, err = cfg.Build()
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
		os.Exit(-2)
	}

	zap.ReplaceGlobals(_logger)
	_sugar = _logger.Sugar()
	_initialized = true
}

func Debug(args ...interface{}) {
	if _initialized {
		_sugar.Debug(args...)
	} else {
		log.Println(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if _initialized {
		_sugar.Debugf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Debugw(msg string, kv ...interface{}) {
	if _initialized {
		_sugar.Debugw(msg, kv...)
	} else {
		log.Print(msg)
		log.Print(kv...)
	}
}

func Info(args ...interface{}) {
	if _initialized {
		_sugar.Info(args...)
	} else {
		log.Println(args...)
	}
}

func Infof(template string, args ...interface{}) {
	if _initialized {
		_sugar.Infof(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Infow(msg string, kv ...interface{}) {
	if _initialized {
		_sugar.Infow(msg, kv...)
	} else {
		log.Print(msg)
		log.Print(kv...)
	}
}

func Warn(args ...interface{}) {
	if _initialized {
		_sugar.Warn(args...)
	} else {
		log.Println(args...)
	}
}

func Warnf(template string, args ...interface{}) {
	if _initialized {
		_sugar.Warnf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Warnw(msg string, kv ...interface{}) {
	if _initialized {
		_sugar.Warnw(msg, kv...)
	} else {
		log.Print(msg)
		log.Print(kv...)
	}
}

func Error(args ...interface{}) {
	if _initialized {
		_sugar.Error(args...)
	} else {
		log.Println(args...)
	}
}

func Errorf(template string, args ...interface{}) {
	if _initialized {
		_sugar.Errorf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Errorw(msg string, kv ...interface{}) {
	if _initialized {
		_sugar.Errorw(msg, kv...)
	} else {
		log.Print(msg)
		log.Print(kv...)
	}
}

func Fatal(args ...interface{}) {
	if _initialized {
		_sugar.Fatal(args...)
	} else {
		log.Println(args...)
	}
}

func Fatalf(template string, args ...interface{}) {
	if _initialized {
		_sugar.Fatalf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Fatalw(msg string, kv ...interface{}) {
	if _initialized {
		_sugar.Fatalw(msg, kv...)
	} else {
		log.Print(msg)
		log.Print(kv...)
	}
}

func Panic(args ...interface{}) {
	if _initialized {
		_sugar.Panic(args...)
	} else {
		log.Println(args...)
	}
}

func Panicf(template string, args ...interface{}) {
	if _initialized {
		_sugar.Panicf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Panicw(msg string, kv ...interface{}) {
	if _initialized {
		_sugar.Panicw(msg, kv...)
	} else {
		log.Print(msg)
		log.Print(kv...)
	}
}

func HasDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
