package logger

import (
	"errors"
	"os"
	"reflect"

	"github.com/AtelierNeon/an-core-api-golang/internal"
)

var (
	providerInstanceType = reflect.TypeOf((*DefaultProvider)(nil)).Elem()
	providerType         = reflect.TypeOf((*Provider)(nil)).Elem()
	instances            = make(map[string]Logger)
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

func GetDefaultProviderOutput() DefaultProviderOutputType {
	return defaultProvider_output
}

func getFullTypeName(t reflect.Type) string {
	if t == nil {
		return ""
	}

	result := ""
	pkgPath := t.PkgPath()
	if pkgPath != "" {
		result += pkgPath + "."
	}
	result += t.Name()
	return result
}

func GetInstance() Logger {
	return GetInstanceWithLabel("")
}

func GetInstanceWithLabel(label string) Logger {
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

func SetDefaultProviderOutput(output DefaultProviderOutputType) {
	defaultProvider_output = output
}
