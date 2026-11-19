package networkmanager

type Provider interface {
	OnGetProtocolStackForLocalAddress(localAddress string) (ProtocolStack, error)
	OnGetProtocolStackForEveryLocalAddress() (ProtocolStack, error)
	OnGetProtocolStackForLocalhost() (ProtocolStack, error)
}
