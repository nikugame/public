package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

//Logger log logic
type Logger struct {
	logger *log.Logger
	file   *os.File
}

//New func(service string) (*Logger,error)
func New(service string) (*Logger, error) {

	var logger *log.Logger
	var file *os.File

	if service != "" {
		filename := fmt.Sprintf("%s_%s.log", service, time.Now().Format("2006_01_02_15_04_05"))
		f, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
		logger = log.New(f, "", log.LstdFlags)
		file = f
	} else {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	l := new(Logger)
	l.logger = logger
	l.file = file

	return l, nil
}

var nikLogger, _ = New("")

//Log func(format string, a ...interface{})
func Log(format string, a ...interface{}) {
	nikLogger.Log(format, a)
}

//Export func(logger *Logger)
func Export(l *Logger) {
	if l != nil {
		nikLogger = l
	}
}

//Close func()
func Close() {
	nikLogger.Close()
}

//Log func(format string, a ...interface{})
func (l *Logger) Log(format string, a ...interface{}) {
	l.logger.Printf(format, a)
}

//Close func()
func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
	l.logger = nil
	l.file = nil
}
