package networkmanager

import "errors"

type DefaultProviderProtocolStack struct {
	ProtocolStack
}

func (dpps *DefaultProviderProtocolStack) CheckPortBindingProcessID(port uint16) (int, error) {
	return 0, errors.New("not implemented yet")
}

func (dpps *DefaultProviderProtocolStack) CheckPortBindingProcessIDList(port uint16) ([]int, error) {
	return nil, errors.New("not implemented yet")
}

func (dpps *DefaultProviderProtocolStack) GetAvailablePort() (uint16, error) {
	return 0, errors.New("not implemented yet")
}

func (dpps *DefaultProviderProtocolStack) IsPortAvailable(port uint16) (bool, error) {
	return false, errors.New("not implemented yet")
}
