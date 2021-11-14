package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger(level string) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			TimeKey:       "time",  //时间对应的key名
			LevelKey:      "level", //日志级别对应的key名
			FunctionKey:   "func",
			CallerKey:     "caller",                      //时间对应的key名
			NameKey:       "path",                        //logger名对应的key名
			StacktraceKey: "stacktrace",                  //栈追踪的key名
			LineEnding:    zapcore.DefaultLineEnding,     //默认换行，即使不设置
			EncodeLevel:   zapcore.LowercaseLevelEncoder, //小写
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder, //记录调用路径格式为package/file:line
			EncodeName:    zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{"stdout", "./log.txt"},
		ErrorOutputPaths: []string{"stderr", "./error.txt"},
	}
	zap.NewProductionEncoderConfig()
	lg, _ = config.Build()
	sugarLogger = lg.Sugar()
}

func GetLogger() *zap.Logger {
	return lg
}

func GetSugarLogger() *zap.SugaredLogger {
	return sugarLogger
}
