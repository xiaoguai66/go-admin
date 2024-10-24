package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zap.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zap.InfoLevel
	}
	//core := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), getStdoutSyncer()), logMode)
	return zap.New(core).Sugar()
}

// {"level":"INFO","time":"2024-10-11 22:03:06","caller":"router/router.go:81","msg":"Server Exist"}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getStdoutSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getWriteSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"

	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"), // megabytes
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"), //days
		Compress:   true,                       // disabled by default
	}

	return zapcore.AddSync(lumberjackSyncer)
}
