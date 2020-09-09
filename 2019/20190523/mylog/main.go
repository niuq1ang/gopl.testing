package main

import (
	"io"
	stdlog "log"
	"os"
	"time"
)

// GlobalLevel 日志输出级别
var GlobalLevel = INFO

// LogLevel 日志级别
type LogLevel byte

const (
	DEBUG LogLevel = iota + 1
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) ToString() string {
	var v string
	switch l {
	case DEBUG:
		v = "Debug"
	case INFO:
		v = "Info"
	case WARN:
		v = "Warn"
	case ERROR:
		v = "Error"
	case FATAL:
		v = "Fatal"
	}
	return v
}

func InitLog() {
	logf, err := os.OpenFile(generateLogFileName(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		stdlog.Fatalln(err)
	}
	stdlog.SetFlags(stdlog.LstdFlags)
	stdlog.SetOutput(io.MultiWriter(logf, os.Stdout))
}

func main() {
	Info("%s%s", "bidfi", "sdfvdf")
	Debug("%s%s", "bidfi", "v打广告")
	Warn("%s%s", "bidfi", "v打广告")
}

func Info(format string, v ...interface{}) {
	printLog(INFO, format, v...)
}

func Debug(format string, v ...interface{}) {
	printLog(DEBUG, format, v...)
}

func Warn(format string, v ...interface{}) {
	printLog(WARN, format, v...)
}
func Error(format string, v ...interface{}) {
	printLog(ERROR, format, v...)
}
func Fatal(format string, v ...interface{}) {
	printLog(FATAL, format, v...)
}

func printLog(level LogLevel, format string, v ...interface{}) {
	if GlobalLevel > level {
		return
	}
	stdlog.SetPrefix(level.ToString() + " ")
	stdlog.Printf(format, v...)
}

func generateLogFileName() string {
	return time.Now().Format("20060203") + ".log"
}
