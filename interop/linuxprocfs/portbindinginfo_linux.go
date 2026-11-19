package linuxprocfs

type PortBindingInfo struct {
	LocalAddress   string
	LocalPort      uint16
	RemoteAddress  string
	RemotePort     uint16
	OwnerProcessID int
	State          string
	Kind           string
	Inode          string
}
