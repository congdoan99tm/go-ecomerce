package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	//sugar := zap.NewExample().Sugar()
	//sugar.Infof("hello name:%s, age:%d", "tipGo", 23)
	//
	//logger := zap.NewExample()
	//logger.Info("Hello", zap.String("name", "TipGo"), zap.Int("age", 23))

	//logger := zap.NewExample()
	//logger.Info("Hello")
	//// Development
	//logger, _ = zap.NewDevelopment()
	//logger.Info("Hello Development")
	//// Production
	//logger, _ = zap.NewProduction()
	//logger.Info("Hello Production")
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile(".log/log.txt", os.O_CREATE|os.O_RDONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
