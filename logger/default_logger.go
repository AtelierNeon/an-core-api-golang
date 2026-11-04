package logger

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/AtelierNeon/an-core-api-golang/internal"
)

type DefaultLogger struct {
	Logger
	provider      Provider
	severity      SeverityType
	severityMutex sync.Mutex
}

func (dl *DefaultLogger) Debug(message string) {
	if dl.GetSeverity() > SeverityTypeDebug {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	err := dl.provider.OnDebug(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.Debug() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) DebugF(format string, args ...any) {
	if dl.GetSeverity() > SeverityTypeDebug {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	message := fmt.Sprintf(format, args...)
	err := dl.provider.OnDebug(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.DebugF() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) Error(message string) {
	if dl.GetSeverity() > SeverityTypeError {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	err := dl.provider.OnError(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.Error() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) ErrorF(format string, args ...any) {
	if dl.GetSeverity() > SeverityTypeError {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	message := fmt.Sprintf(format, args...)
	err := dl.provider.OnError(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.ErrorF() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) Fatal(message string) {
	if dl.GetSeverity() > SeverityTypeFatal {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	err := dl.provider.OnFatal(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.Fatal() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) FatalF(format string, args ...any) {
	if dl.GetSeverity() > SeverityTypeFatal {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	message := fmt.Sprintf(format, args...)
	err := dl.provider.OnFatal(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.FatalF() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) GetSeverity() SeverityType {
	dl.severityMutex.Lock()
	defer dl.severityMutex.Unlock()

	return dl.severity
}

func (dl *DefaultLogger) Info(message string) {
	if dl.GetSeverity() > SeverityTypeInfo {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	err := dl.provider.OnInfo(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.Info() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) InfoF(format string, args ...any) {
	if dl.GetSeverity() > SeverityTypeInfo {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	message := fmt.Sprintf(format, args...)
	err := dl.provider.OnInfo(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.InfoF() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) Warn(message string) {
	if dl.GetSeverity() > SeverityTypeWarn {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	err := dl.provider.OnWarn(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.Warn() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) WarnF(format string, args ...any) {
	if dl.GetSeverity() > SeverityTypeWarn {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	message := fmt.Sprintf(format, args...)
	err := dl.provider.OnWarn(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.WarnF() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) SetSeverity(s SeverityType) Logger {
	dl.severityMutex.Lock()
	defer dl.severityMutex.Unlock()

	dl.severity = s
	return dl
}

func (dl *DefaultLogger) Trace(message string) {
	if dl.GetSeverity() > SeverityTypeTrace {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	err := dl.provider.OnTrace(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.Trace() %s\n",
			err.Error(),
		)
	}
}

func (dl *DefaultLogger) TraceF(format string, args ...any) {
	if dl.GetSeverity() > SeverityTypeTrace {
		return
	}

	tag := ""
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		if showPackageName {
			detail := runtime.FuncForPC(pc)
			if detail != nil {
				tag = detail.Name()
			}
		} else {
			tag = fmt.Sprintf("%s:%d", file, line)
		}
	}

	message := fmt.Sprintf(format, args...)
	err := dl.provider.OnTrace(
		tag,
		message,
	)
	if err != nil {
		internal.Fprintf(
			os.Stderr,
			"[Fatal] logger.DefaultLogger.TraceF() %s\n",
			err.Error(),
		)
	}
}
