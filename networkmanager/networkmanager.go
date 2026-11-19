package networkmanager

type NetworkManager interface {
	GetProtocolStackForLocalAddress(localAddress string) (ProtocolStack, error)
	GetProtocolStackForEveryLocalAddress() (ProtocolStack, error)
	GetProtocolStackForLocalhost() (ProtocolStack, error)
}
