package logger

import (
	"errors"
	"os"
	"reflect"
	"sync"

	"github.com/AtelierNeon/an-core-api-golang/internal"
)

var (
	providerInstanceType = reflect.TypeOf((*DefaultProvider)(nil)).Elem()
	providerType         = reflect.TypeOf((*Provider)(nil)).Elem()
	instances            = make(map[string]Logger)
	instancesMutex       = &sync.Mutex{}
	severity             = SeverityTypeDebug
	showPackageName      = false
)

func doGetInstanceWithLabel(
	t reflect.Type,
	label string) (Logger, error) {
	if t == nil {
		return nil, errors.New("reflect.Type is not found")
	}

	providerFullTypeName := getFullTypeName(t)
	key := providerFullTypeName + "_" + label
	if instance, ok := instances[key]; ok {
		return instance, nil
	}

	internal.Fprintf(
		os.Stderr,
		"Initializing %s...\n",
		key,
	)

	providerInstance := reflect.New(t).Interface().(Provider)
	instance := &DefaultLogger{
		provider: providerInstance,
		severity: severity,
	}

	instances[key] = instance
	return instance, nil
}

func GetDefaultProviderOutputType() DefaultProviderOutputType {
	defaultProvider_outputTypeMutex.Lock()
	defer defaultProvider_outputTypeMutex.Unlock()

	return defaultProvider_outputType
}

func getFullTypeName(t reflect.Type) string {
	if t == nil {
		return ""
	}

	for t.Kind() == reflect.Pointer && t.Name() == "" {
		t = t.Elem()
	}

	result := ""
	pkgPath := t.PkgPath()
	if pkgPath != "" {
		result = pkgPath + "."
	}
	result += t.Name()
	return result
}

func GetInstance() Logger {
	return GetInstanceWithLabel("")
}

func GetInstanceWithLabel(label string) Logger {
	instancesMutex.Lock()
	defer instancesMutex.Unlock()

	instance, err := doGetInstanceWithLabel(
		providerInstanceType,
		label,
	)
	if err != nil {
		panic("Can not get instance")
	}

	return instance
}

func GetInstanceWithType(t reflect.Type) Logger {
	fullTypeName := getFullTypeName(t)
	if fullTypeName != "" {
		return GetInstanceWithLabel(fullTypeName)
	}

	return GetInstance()
}

func Register(t reflect.Type) {
	if t == nil || t.Kind() == reflect.Pointer ||
		!reflect.PointerTo(t).Implements(providerType) {
		panic("registered type must be a non-pointer type whose pointer implements Provider")
	}

	instancesMutex.Lock()
	defer instancesMutex.Unlock()

	if providerInstanceType == t {
		return
	}

	providerInstanceType = t
	internal.Fprintf(
		os.Stderr,
		"Registered default %s type to %s\n",
		reflect.TypeOf((*Provider)(nil)).Elem().Name(),
		providerInstanceType.Name(),
	)
}

func SetDefaultProviderOutputType(output DefaultProviderOutputType) {
	defaultProvider_outputTypeMutex.Lock()
	defer defaultProvider_outputTypeMutex.Unlock()

	defaultProvider_outputType = output
}
