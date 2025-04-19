package engine

import "bufio"

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUB
	INFO
	WARNING
	ERROR
	FATAL
)	

type logger struct {
	level LogLevel
	writer bufio.Writer
}

var log *logger

func InitLogger(minLogLevel LogLevel, writer bufio.Writer) {
	if log != nil {
		panic("Logger already initialized")
	}

	log = &logger{
		level: minLogLevel,
		writer: writer,
	}
}

func Log( level LogLevel, message string) {
	if log != nil && level >= log.level {
		log.writer.WriteString(message)
		log.writer.Flush()
	}
}

func LogTrace(message string) {
	Log(TRACE, message)
}

func LogDebug(message string) {
	Log(DEBUB, message)
}

func LogInfo(message string) {
	Log(INFO, message)
}

func LogWarning(message string) {
	Log(WARNING, message)
}

func LogError(message string) {
	Log(ERROR, message)
}

func LogFatal(message string) {
	Log(FATAL, message)
	panic(message)
}
