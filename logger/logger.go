package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

// daywiseWriter is a custom WriteSyncer that switches log files daily.
type daywiseWriter struct {
	mu         sync.Mutex
	currentDay int
	logFile    *os.File
}

func newDaywiseWriter(logFilePath string) *daywiseWriter {
	writer := &daywiseWriter{}
	writer.rotateLogFile(logFilePath)
	return writer
}

func (w *daywiseWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Check if the day has changed
	today := time.Now().Day()
	if today != w.currentDay {
		w.rotateLogFile("./logs/debug.log")
	}

	return w.logFile.Write(p)
}

func (w *daywiseWriter) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.logFile.Sync()
}

func (w *daywiseWriter) rotateLogFile(logFilePath string) {
	if w.logFile != nil {
		w.logFile.Close()
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	w.logFile = file
	w.currentDay = time.Now().Day()
}

func init() {
	logConfig := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(logConfig.EncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(newDaywiseWriter("./logs/debug.log"))),
		logConfig.Level,
	)

	Log = zap.New(core)
}

// Info logs an informational message.
func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

// Error logs an error message.
func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	Log.Error(msg, tags...)
	Log.Sync()
}
