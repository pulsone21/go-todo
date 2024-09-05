package logg

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

type LogLevel int

const (
	DEBUG LogLevel = 3
	INFO  LogLevel = 2
	WARN  LogLevel = 1
	ERROR LogLevel = 0
)

var (
	def      = log.Default()
	loglevel = 0
)

func SetOutput(logPath string, maxSize, level int) {
	def.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    maxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	})
	loglevel = level
	Debug("logger setup, logpath: &v, logSize: &v, logLevel: &b", logPath, maxSize, level)
}

func Debug(msg string, a ...any) {
	if loglevel > 2 {
		log.Printf("DEBUG %v\n", msg)
	}
}

func Info(msg string, a ...any) {
	if loglevel >= 2 {
		log.Printf("INFO %v\n", msg)
	}
}

func Warn(msg string, a ...any) {
	if loglevel >= 1 {
		log.Printf("WARN %v\n", msg)
	}
}

func Error(msg string, a ...any) {
	log.Printf("ERROR %v\n", msg)
}

func Fatal(msg string, a ...any) {
	log.Fatalf("FATAL %v\n", msg)
}
