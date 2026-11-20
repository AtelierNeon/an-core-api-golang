package darwinlibproc

import "github.com/AtelierNeon/an-core-api-golang/interop"

type PortBindingInfo struct {
	LocalAddress   string
	LocalPort      uint16
	RemoteAddress  string
	RemotePort     uint16
	OwnerProcessID interop.ProcessID
	State          interop.TcpState
	Kind           interop.PortBindingKind
}
