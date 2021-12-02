package logger

import (
	"os"

	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLogger(isProduction bool) *zap.Logger {
	var zapCfg zap.Config
	if isProduction {
		zapCfg = zap.NewProductionConfig()
	} else {
		zapCfg = zap.NewDevelopmentConfig()
	}
	zapCfg.DisableStacktrace = true
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapLogger, _ := zapCfg.Build()
	zap.ReplaceGlobals(zapLogger)
	return zapLogger
}

func AddGrayLog(serviceName, grayAddress string, l *zap.Logger) (*zap.Logger, error) {
	grayCore, err := gelf.NewCore(
		gelf.Addr(grayAddress),
		gelf.Host(serviceName),
	)
	if err != nil {
		return nil, err
	}
	wrappedLog := attachCoreToLogger(grayCore, l)
	zap.ReplaceGlobals(wrappedLog)
	return wrappedLog, nil
}

func InitZapFileConsole(debug bool, f *os.File) *zap.Logger {
	prC := zap.NewProductionEncoderConfig()
	prC.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(prC)

	dvC := zap.NewDevelopmentEncoderConfig()
	dvC.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(dvC)

	level := zap.InfoLevel
	if debug {
		level = zap.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	l := zap.New(core)

	return l
}

func attachCoreToLogger(newCore zapcore.Core, l *zap.Logger) *zap.Logger {
	return l.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, newCore)
	}))
}
