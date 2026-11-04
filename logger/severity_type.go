package logger

type SeverityType int32

const (
	SeverityTypeTrace SeverityType = iota
	SeverityTypeDebug
	SeverityTypeInfo
	SeverityTypeWarn
	SeverityTypeError
	SeverityTypeFatal
)
