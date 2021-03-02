package plog

import (
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
)

type LogZap struct {
	logger *zap.SugaredLogger
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *LogZap) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *LogZap) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *LogZap) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *LogZap) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *LogZap) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *LogZap) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *LogZap) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *LogZap) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *LogZap) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *LogZap) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *LogZap) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *LogZap) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *LogZap) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *LogZap) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (l *LogZap) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *LogZap) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *LogZap) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *LogZap) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (l *LogZap) DPanicw(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (l *LogZap) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (l *LogZap) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues)
}

var localLog *LogZap

func init() {
	localLog = nil
}

func createLog() *LogZap {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Error to create logger...")
	}

	logZap := LogZap{logger: logger.Sugar()}
	return &logZap
}

func GetLog() *LogZap {
	if localLog == nil {
		localLog = createLog()
	}

	return localLog
}

func Provide(container *dig.Container) {
	container.Provide(func() Log {
		return GetLog()
	})
}
