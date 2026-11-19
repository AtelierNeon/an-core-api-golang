package interop

type UdpTableClass uint32

const (
	UdpTableClassBasic       UdpTableClass = /* UDP_TABLE_BASIC        */ 0
	UdpTableClassOwnerPid    UdpTableClass = /* UDP_TABLE_OWNER_PID    */ 1
	UdpTableClassOwnerModule UdpTableClass = /* UDP_TABLE_OWNER_MODULE */ 2
)
