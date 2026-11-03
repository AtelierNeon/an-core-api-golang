package logger

import (
	"fmt"
	"os"

	"github.com/AtelierNeon/an-core-api-golang/internal"
)

type DefaultProvider struct {
	Provider
}

var (
	defaultProvider_output DefaultProviderOutputType = DefaultProviderOutputTypeStandardOutput
)

func doPrint(
	level string,
	tag string,
	message string) error {
	line := fmt.Sprintf("[%s][%s] %s", level, tag, message)

	switch defaultProvider_output {
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

func (cp *DefaultProvider) OnDebug(tag string, message string) error {
	return doPrint(
		"Debug",
		tag,
		message,
	)
}

func (cp *DefaultProvider) OnError(tag string, message string) error {
	return doPrint(
		"Error",
		tag,
		message,
	)
}

func (cp *DefaultProvider) OnFatal(tag string, message string) error {
	return doPrint(
		"Fatal",
		tag,
		message,
	)
}

func (cp *DefaultProvider) OnInfo(tag string, message string) error {
	return doPrint(
		"Info",
		tag,
		message,
	)
}

func (cp *DefaultProvider) OnWarn(tag string, message string) error {
	return doPrint(
		"Warn",
		tag,
		message,
	)
}

func (cp *DefaultProvider) OnTrace(tag string, message string) error {
	return doPrint(
		"Trace",
		tag,
		message,
	)
}
