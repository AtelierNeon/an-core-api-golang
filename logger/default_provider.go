package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/AtelierNeon/an-core-api-golang/internal"
)

type DefaultProvider struct {
	Provider
}

var (
	defaultProvider_outputType      DefaultProviderOutputType = DefaultProviderOutputTypeStandardOutput
	defaultProvider_outputTypeMutex                           = &sync.Mutex{}
)

func doPrint(
	level string,
	tag string,
	message string) error {
	line := fmt.Sprintf("[%s][%s] %s", level, tag, message)

	outputType := GetDefaultProviderOutputType()

	switch outputType {
	case DefaultProviderOutputTypeStandardOutput:
		internal.Fprintln(os.Stdout, line)
		return nil
	case DefaultProviderOutputTypeStandardError:
		internal.Fprintln(os.Stderr, line)
		return nil
	case DefaultProviderOutputTypeSuppressed:
		return nil
	}

	fmt.Println(line)
	return nil
}

func (dp *DefaultProvider) OnDebug(tag string, message string) error {
	return doPrint(
		"Debug",
		tag,
		message,
	)
}

func (dp *DefaultProvider) OnError(tag string, message string) error {
	return doPrint(
		"Error",
		tag,
		message,
	)
}

func (dp *DefaultProvider) OnFatal(tag string, message string) error {
	return doPrint(
		"Fatal",
		tag,
		message,
	)
}

func (dp *DefaultProvider) OnInfo(tag string, message string) error {
	return doPrint(
		"Info",
		tag,
		message,
	)
}

func (dp *DefaultProvider) OnWarn(tag string, message string) error {
	return doPrint(
		"Warn",
		tag,
		message,
	)
}

func (dp *DefaultProvider) OnTrace(tag string, message string) error {
	return doPrint(
		"Trace",
		tag,
		message,
	)
}
