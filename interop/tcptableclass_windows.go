package interop

type TcpTableClass WindowsEnum

const (
	TcpTableClassBasicListener          TcpTableClass = /* TCP_TABLE_BASIC_LISTENER           */ 0
	TcpTableClassBasicConnections       TcpTableClass = /* TCP_TABLE_BASIC_CONNECTIONS        */ 1
	TcpTableClassBasicAll               TcpTableClass = /* TCP_TABLE_BASIC_ALL                */ 2
	TcpTableClassOwnerPidListener       TcpTableClass = /* TCP_TABLE_OWNER_PID_LISTENER       */ 3
	TcpTableClassOwnerPidConnections    TcpTableClass = /* TCP_TABLE_OWNER_PID_CONNECTIONS    */ 4
	TcpTableClassOwnerPidAll            TcpTableClass = /* TCP_TABLE_OWNER_PID_ALL            */ 5
	TcpTableClassOwnerModuleListener    TcpTableClass = /* TCP_TABLE_OWNER_MODULE_LISTENER    */ 6
	TcpTableClassOwnerModuleConnections TcpTableClass = /* TCP_TABLE_OWNER_MODULE_CONNECTIONS */ 7
	TcpTableClassOwnerModuleAll         TcpTableClass = /* TCP_TABLE_OWNER_MODULE_ALL         */ 8
)
