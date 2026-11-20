package networkmanager

type ProtocolStack interface {
	CheckPortBindingProcessIDList(port uint16) ([]int, error)
	GetAvailablePort() (uint16, error)
	IsPortAvailable(port uint16) (bool, error)
}
