package prettylog

import (
	"io"
	"log/slog"
	"os"
)

type severity int

func (s severity) String() string {
	switch s {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}

	return "UNKNOWN"
}

const (
	INFO severity = iota
	WARN
	ERROR
)

type logger struct {
	w io.Writer
}

func (l *logger) prepare(sev severity, log string, keyVals ...any) {
	var out []byte
	out = append(out)

}

func (l *logger) Info(log string, logs ...any) {
	slog.Info(log, logs...)
	l.prepare(INFO, log, logs...)
}

func New() *logger {
	return &logger{os.Stdout}
}
