package networkmanager

type DefaultProvider struct {
	Provider
}

func (dp *DefaultProvider) OnGetProtocolStackForLocalAddress(localAddress string) (ProtocolStack, error) {
	return &DefaultProviderProtocolStack{}, nil
}

func (dp *DefaultProvider) OnGetProtocolStackForEveryLocalAddress() (ProtocolStack, error) {
	return &DefaultProviderProtocolStack{}, nil
}

func (dp *DefaultProvider) OnGetProtocolStackForLocalhost() (ProtocolStack, error) {
	return &DefaultProviderProtocolStack{}, nil
}
