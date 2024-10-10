package log

import (
	"regexp"
)

type Logger struct {
	LogLevel     LogLevel
	IsWriteFile  bool
	LogFilePath  *string
	VarsRepRegex *regexp.Regexp
}

func (l *Logger) InitLogFile() {
	if !l.IsWriteFile {

	}
}
