package networkmanager

type ProtocolStack interface {
	CheckPortBindingProcessID(port uint16) (int, error)
	GetAvailablePort() (uint16, error)
	IsPortAvailable(port uint16) (bool, error)
}
