package interop

type PortBindingKind uint32

const (
	PortBindingKindUnknown PortBindingKind = iota
	PortBindingKindTCP4
	PortBindingKindUDP4
	PortBindingKindTCP6
	PortBindingKindUDP6
)
