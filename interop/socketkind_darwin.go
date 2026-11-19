package interop

type SocketKind DarwinInt

/* https://github.com/apple/darwin-xnu/blob/main/bsd/sys/proc_info.h */
const (
	SocketKindGeneric   SocketKind = /* SOCKINFO_GENERIC    */ 0
	SocketKindIn        SocketKind = /* SOCKINFO_IN         */ 1
	SocketKindTcp       SocketKind = /* SOCKINFO_TCP        */ 2
	SocketKindUn        SocketKind = /* SOCKINFO_UN         */ 3
	SocketKindNdrv      SocketKind = /* SOCKINFO_NDRV       */ 4
	SocketKindKernEvent SocketKind = /* SOCKINFO_KERN_EVENT */ 5
	SocketKindKernCtl   SocketKind = /* SOCKINFO_KERN_CTL   */ 6
	SocketKindVsock     SocketKind = /* SOCKINFO_VSOCK      */ 7
)
