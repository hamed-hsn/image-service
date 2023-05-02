package simplelogger

import (
	"go.uber.org/zap"
)

type Logger struct {
	zLogger *zap.Logger
	sugar   *zap.SugaredLogger
}

func (l Logger) Error(msg string, keyValue ...interface{}) {
	l.sugar.Errorw(msg, keyValue...)
}

func (l Logger) Info(msg string, keyValue ...interface{}) {

	l.sugar.Infow(msg, keyValue...)
}

func (l Logger) Warning(msg string, keyValue ...interface{}) {
	l.sugar.Warnw(msg, keyValue...)
}

func (l Logger) Debug(msg string, keyValue ...interface{}) {
	l.sugar.Debugw(msg, keyValue...)
}

func New() *Logger {
	zlogger, _ := zap.NewProduction()
	sugar := zlogger.Sugar()

	return &Logger{sugar: sugar}
}
