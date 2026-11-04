package desktopadapter

import (
	"reflect"

	"github.com/AtelierNeon/an-core-api-golang/internal"
	"github.com/AtelierNeon/an-core-api-golang/logger"
)

var (
	loggerProviderInstanceType = reflect.TypeFor[LogrusProvider]()
)

func Configure() bool {
	return ConfigureWithConfig(NewConfig())
}

func ConfigureWithConfig(config *Config) bool {
	realConfig := config
	if realConfig == nil {
		realConfig = NewConfig()
	}

	internal.DisableLogging(realConfig.DisableInternalLogger)

	logger.SetDefaultProviderOutputType(logger.DefaultProviderOutputTypeSuppressed)
	logger.Register(loggerProviderInstanceType)

	return true
}

func NewConfig() *Config {
	return &Config{}
}
