package log

type LogLevel int

const (
	OFF LogLevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)
