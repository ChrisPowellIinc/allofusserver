package log

import (
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func Register(logDir string) {
	log = logrus.New()
	log.AddHook(ContextHook{})

	// Rotate log files if the log file grows up to 50mb in size
	// (Whichever comes first)
	// Backups use the log file name given to Logger,
	// in the form name-timestamp.ext.
	// Eg for /var/log/foo/server.log, a backup created at 6:30pm
	// on Nov 11 2016 would use the filename
	// /var/log/foo/server-2016-11-04T18-30-00.000.log
	fileLogger := &lumberjack.Logger{
		Filename:   path.Join(logDir, "server.log"),
		MaxSize:    50, // megabytes
		MaxBackups: 20,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	}

	// Log to both stdout and to filesystem
	logWriter := io.MultiWriter(os.Stdout, fileLogger)
	log.Formatter = &logrus.TextFormatter{}
	log.Out = logWriter
}

func Log() *logrus.Logger {
	return log
}

func Info(infoData ...interface{}) {
	log.Info(infoData...)
}

func Debug(debugData interface{}) {
	log.Debug(debugData)
}

func Error(err ...interface{}) {
	if err == nil || err[0] == nil {
		return
	}
	log.Error(err...)
}

func Warn(warnData ...interface{}) {
	log.Warn(warnData...)
}
func Fatal(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Aliases
var Panicf = Fatalf
var Fatalf = Fatal
var Panic = Fatal

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Println(args ...interface{}) {
	log.Println(args...)
}

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(8); ok {
		funcName := runtime.FuncForPC(pc).Name()

		entry.Data["file"] = path.Base(file)
		entry.Data["func"] = path.Base(funcName)
		entry.Data["line"] = line
	}

	return nil
}
