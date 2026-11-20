package desktopadapter

import "github.com/AtelierNeon/an-core-api-golang/networkmanager"

type NetworkManagerProvider struct {
	networkmanager.Provider
}

func (nmp *NetworkManagerProvider) OnGetProtocolStackForLocalAddress(localAddress string) (networkmanager.ProtocolStack, error) {
	return &NetworkManagerProviderProtocolStack{
		localAddress: localAddress,
	}, nil
}

func (nmp *NetworkManagerProvider) OnGetProtocolStackForEveryLocalAddress() (networkmanager.ProtocolStack, error) {
	return nmp.OnGetProtocolStackForLocalAddress("")
}

func (nmp *NetworkManagerProvider) OnGetProtocolStackForLocalhost() (networkmanager.ProtocolStack, error) {
	return nmp.OnGetProtocolStackForLocalAddress("localhost")
}
