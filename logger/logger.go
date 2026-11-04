package logger

type Logger interface {
	Debug(message string)
	DebugF(format string, args ...any)
	Error(message string)
	ErrorF(format string, args ...any)
	Fatal(message string)
	FatalF(format string, args ...any)
	GetSeverity() SeverityType
	Info(message string)
	InfoF(format string, args ...any)
	SetSeverity(s SeverityType) Logger
	Warn(message string)
	WarnF(format string, args ...any)
	Trace(message string)
	TraceF(format string, args ...any)
}
