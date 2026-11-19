package interop

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type ProcessFdType uint32

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
const (
	ProcessFdTypeAtalk     ProcessFdType = /* PROX_FDTYPE_ATALK     */ 0
	ProcessFdTypeVnode     ProcessFdType = /* PROX_FDTYPE_VNODE     */ 1
	ProcessFdTypeSocket    ProcessFdType = /* PROX_FDTYPE_SOCKET    */ 2
	ProcessFdTypePshm      ProcessFdType = /* PROX_FDTYPE_PSHM      */ 3
	ProcessFdTypePsem      ProcessFdType = /* PROX_FDTYPE_PSEM      */ 4
	ProcessFdTypeKqueue    ProcessFdType = /* PROX_FDTYPE_KQUEUE    */ 5
	ProcessFdTypePipe      ProcessFdType = /* PROX_FDTYPE_PIPE      */ 6
	ProcessFdTypeFsEvents  ProcessFdType = /* PROX_FDTYPE_FSEVENTS  */ 7
	ProcessFdTypeNetPolicy ProcessFdType = /* PROX_FDTYPE_NETPOLICY */ 9
)
