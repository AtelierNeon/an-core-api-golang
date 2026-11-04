package desktopadapter

import (
	"fmt"
	"sync"

	"github.com/AtelierNeon/an-core-api-golang/logger"
	"github.com/sirupsen/logrus"
)

type LogrusProvider struct {
	logger.Provider
	logrusLogger    *logrus.Logger
	initializedOnce sync.Once
}

func (lp *LogrusProvider) initialize() {
	lp.initializedOnce.Do(func() {
		if lp.logrusLogger == nil {
			lp.logrusLogger = logrus.New()
			lp.logrusLogger.SetFormatter(&logrus.JSONFormatter{})
			lp.logrusLogger.SetLevel(logrus.TraceLevel)
		}
	})
}

func (lp *LogrusProvider) OnDebug(tag string, message string) error {
	lp.initialize()

	line := fmt.Sprintf(" %s() %s", tag, message)
	lp.logrusLogger.Log(logrus.DebugLevel, line)
	return nil
}

func (lp *LogrusProvider) OnError(tag string, message string) error {
	lp.initialize()

	line := fmt.Sprintf(" %s() %s", tag, message)
	lp.logrusLogger.Log(logrus.ErrorLevel, line)
	return nil
}

func (lp *LogrusProvider) OnFatal(tag string, message string) error {
	lp.initialize()

	line := fmt.Sprintf(" %s() %s", tag, message)
	lp.logrusLogger.Log(logrus.FatalLevel, line)
	return nil
}

func (lp *LogrusProvider) OnInfo(tag string, message string) error {
	lp.initialize()

	line := fmt.Sprintf(" %s() %s", tag, message)
	lp.logrusLogger.Log(logrus.InfoLevel, line)
	return nil
}

func (lp *LogrusProvider) OnWarn(tag string, message string) error {
	lp.initialize()

	line := fmt.Sprintf(" %s() %s", tag, message)
	lp.logrusLogger.Log(logrus.WarnLevel, line)
	return nil
}

func (lp *LogrusProvider) OnTrace(tag string, message string) error {
	lp.initialize()

	line := fmt.Sprintf(" %s() %s", tag, message)
	lp.logrusLogger.Log(logrus.TraceLevel, line)
	return nil
}
