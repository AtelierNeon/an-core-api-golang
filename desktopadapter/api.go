package desktopadapter

import (
	"reflect"

	"github.com/AtelierNeon/an-core-api-golang/internal"
	"github.com/AtelierNeon/an-core-api-golang/logger"
	"github.com/AtelierNeon/an-core-api-golang/processmanager"
)

var (
	loggerProviderInstanceType         = reflect.TypeFor[LogrusProvider]()
	processManagerProviderInstanceType = reflect.TypeFor[ProcessManagerProvider]()
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

	logger.Register(loggerProviderInstanceType)
	processmanager.Register(processManagerProviderInstanceType)

	return true
}

func NewConfig() *Config {
	return &Config{}
}
