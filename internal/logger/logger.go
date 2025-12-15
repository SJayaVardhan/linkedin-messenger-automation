package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Init() {
	l, _ := zap.NewProduction()
	Log = l.Sugar()
}
