package networkmanager

type DefaultNetworkManager struct {
	NetworkManager
	provider Provider
}

func (dnm *DefaultNetworkManager) GetProtocolStackForLocalAddress(localAddress string) (ProtocolStack, error) {
	return dnm.provider.OnGetProtocolStackForLocalAddress(localAddress)
}

func (dnm *DefaultNetworkManager) GetProtocolStackForEveryLocalAddress() (ProtocolStack, error) {
	return dnm.provider.OnGetProtocolStackForEveryLocalAddress()
}

func (dnm *DefaultNetworkManager) GetProtocolStackForLocalhost() (ProtocolStack, error) {
	return dnm.provider.OnGetProtocolStackForLocalhost()
}
