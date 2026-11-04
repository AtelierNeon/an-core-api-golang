package logger

type Provider interface {
	OnDebug(tag string, message string) error
	OnError(tag string, message string) error
	OnFatal(tag string, message string) error
	OnInfo(tag string, message string) error
	OnWarn(tag string, message string) error
	OnTrace(tag string, message string) error
}
