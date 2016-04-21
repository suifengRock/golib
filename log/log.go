package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// terminal color
var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// log level
const (
	NOTSET   = 0
	DEBUG    = 10
	INFO     = 20
	WARNING  = 30
	ERROR    = 40
	CRITICAL = 50
)

func getLevelStrAndColor(level uint32) (string, string) {
	switch level {
	case DEBUG:
		return "[Debug]", magenta
	case INFO:
		return "[INFO]", green
	case WARNING:
		return "[Warning]", yellow
	case ERROR:
		return "[Error]", red
	case CRITICAL:
		return "[Critical]", cyan
	default:
		return "", ""
	}
}

type Logger struct {
	*log.Logger
	enableColor bool
	level       uint32
	prefix      string
	calldepth   int
}

func New(out io.Writer, prefix string, color bool, level uint32, flag int) *Logger {
	l := &Logger{enableColor: color, level: level, prefix: prefix}
	l.calldepth = 4
	l.Logger = log.New(out, prefix, flag)
	return l
}

func (l *Logger) SetCalldepth(calldepth int) {
	l.calldepth = calldepth
}

func (l *Logger) SetLevel(level uint32) {
	l.level = level
}

func (l *Logger) SetEnableColor(enableColor bool) {
	l.enableColor = enableColor
}

func (l *Logger) content(level uint32, v string, args ...interface{}) {
	levelStr, levelColor := getLevelStrAndColor(level)
	var content string
	var prefix string
	if l.enableColor {
		prefix = fmt.Sprintf("%s%s %s\t", levelColor, l.prefix, levelStr)
		content = fmt.Sprintf("%s%s", reset, fmt.Sprintf(v, args...))
	} else {
		prefix = fmt.Sprintf("%s %s\t", l.prefix, levelStr)
		content = fmt.Sprintf("%s", fmt.Sprintf(v, args...))

	}
	l.SetPrefix(prefix)
	l.print(content)
}
func (l *Logger) print(v string) {
	l.Output(l.calldepth, v)
}

func (l *Logger) Debug(v string, args ...interface{}) {
	if l.level > DEBUG {
		return
	}
	l.content(DEBUG, v, args...)
}

func (l *Logger) Info(v string, args ...interface{}) {
	if l.level > INFO {
		return
	}
	l.content(INFO, v, args...)
}

func (l *Logger) Warning(v string, args ...interface{}) {
	if l.level > WARNING {
		return
	}
	l.content(WARNING, v, args...)
}

func (l *Logger) Error(v string, args ...interface{}) {
	if l.level > ERROR {
		return
	}
	l.content(ERROR, v, args...)
}

func (l *Logger) Fatal(v string, args ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	l.content(CRITICAL, v, args...)
	panic("log fatal")
}

var logger = New(os.Stderr, "[App]", true, DEBUG, log.Lshortfile|log.LstdFlags)

func init() {
	logger.SetCalldepth(5)
}

func SetEnableColor(enableColor bool) {
	logger.SetEnableColor(enableColor)
}

func Debug(v string, args ...interface{}) {

	logger.Debug(v, args...)
}

func Info(v string, args ...interface{}) {
	logger.Info(v, args...)
}

func Warning(v string, args ...interface{}) {
	logger.Warning(v, args...)
}

func Error(v string, args ...interface{}) {
	logger.Error(v, args...)
}

func Fatal(v string, args ...interface{}) {
	logger.Fatal(v, args...)
}
