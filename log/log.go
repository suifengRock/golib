package log

import (
    "log"
    "os"
    "io"
)


type Logger struct{
    *log.Logger
    debug bool
    prefix string
}

func New(out io.Writer, prefix string, debug bool, flag int) *Logger {
    l := &Logger{debug: debug, prefix: prefix}
    l.Logger = log.New(out, "", flag)
    return l
}

func (l *Logger) print(v string) {
    if !l.debug {
        return 
    }
    l.Output(3, v)
}

func (l *Logger) Trace(v string) {
   
    str := "[Trace] "+l.prefix+v
    l.print(str)
}

func (l *Logger) Info(v string) {
    
    str := "[Info] "+l.prefix+v
    l.print(str)
}

func (l *Logger) Warning(v string) {
    str := "[Warning] "+l.prefix+v
    l.print(str)
}

func (l *Logger) Fatal(v string) {
    str := "[Fatal] "+l.prefix+v
    l.print(str)
}



var logger = New(os.Stderr, "", true, log.Lshortfile | log.LstdFlags)

func Trace(v string) {
    str := "[Trace] "+logger.prefix+v
    logger.print(str)
}

func Info(v string) {
    str := "[Info] "+logger.prefix+v
    logger.print(str)
}

func Warning(v string) {
    str := "[Warning] "+logger.prefix+v
    logger.print(str)
}

func Fatal(v string) {
    str := "[Fatal] "+logger.prefix+v
    logger.print(str)
}


