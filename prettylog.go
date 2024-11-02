package prettylog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

type severity int

func (s severity) String() string {
	switch s {
	case INFO:
		return prettyString("INFO", BLUE)
	case DEBUG:
		return prettyString("DEBUG", GREEN)
	case WARN:
		return prettyString("WARN", ORANGE)
	case ERROR:
		return prettyString("ERROR", RED)
	}

	return "UNKNOWN"
}

func (s severity) color() string {
	switch s {
	case INFO:
		return BLUE
	case DEBUG:
		return GREEN
	case WARN:
		return ORANGE
	case ERROR:
		return RED
	}

	return ""
}

const (
	INFO severity = iota
	DEBUG
	WARN
	ERROR

	BLUE   = "\x1b[36m"
	GREEN  = "\x1b[32m"
	ORANGE = "\x1b[33m"
	RED    = "\x1b[31m"
	RESET  = "\x1b[0m"
)

type Logger struct {
	w io.Writer
}

func (l *Logger) print(out string) {
	log.Println(out)
}

func (l *Logger) prepare(sev severity, msg string) string {
	out := strings.Builder{}

	out.WriteString(sev.String())

	out.WriteString(" ")
	if sev == DEBUG {
		pc, _, _, _ := runtime.Caller(2)
		funcName := runtime.FuncForPC(pc).Name()

		out.WriteString(prettyString("["+funcName+"]", sev.color()) + " ")
	}
	out.WriteString(msg)

	return out.String()
}

func (l *Logger) Info(log string, args ...any) {
	str := l.prepare(INFO, log)
	l.print(str)
}

func (l *Logger) Infof(log string, args ...any) {
	str := l.prepare(INFO, log)
	l.print(fmt.Sprintf(str, args...))
}

func (l *Logger) Debug(log string) {
	str := l.prepare(DEBUG, log)
	l.print(str)
}

func (l *Logger) Debugf(log string, args ...any) {
	str := l.prepare(DEBUG, log)
	l.print(fmt.Sprintf(str, args...))
}

func (l *Logger) Warn(log string) {
	str := l.prepare(WARN, log)
	l.print(str)
}

func (l *Logger) Warnf(log string, args ...any) {
	str := l.prepare(WARN, log)
	l.print(fmt.Sprintf(str, args...))
}

func (l *Logger) Error(log string) {
	str := l.prepare(ERROR, log)
	l.print(str)
}

func (l *Logger) Errorf(log string, args ...any) {
	str := l.prepare(ERROR, log)
	l.print(fmt.Sprintf(str, args...))
}

func New() *Logger {
	return &Logger{os.Stdout}
}

func prettyString(str, color string) string {
	return color + str + RESET
}
