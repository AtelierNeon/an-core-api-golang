package interop

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type ProcessFDType uint32

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
const (
	ProcessFDTypeAtalk     ProcessFDType = /* PROX_FDTYPE_ATALK     */ 0
	ProcessFDTypeVnode     ProcessFDType = /* PROX_FDTYPE_VNODE     */ 1
	ProcessFDTypeSocket    ProcessFDType = /* PROX_FDTYPE_SOCKET    */ 2
	ProcessFDTypePshm      ProcessFDType = /* PROX_FDTYPE_PSHM      */ 3
	ProcessFDTypePsem      ProcessFDType = /* PROX_FDTYPE_PSEM      */ 4
	ProcessFDTypeKqueue    ProcessFDType = /* PROX_FDTYPE_KQUEUE    */ 5
	ProcessFDTypePipe      ProcessFDType = /* PROX_FDTYPE_PIPE      */ 6
	ProcessFDTypeFsEvents  ProcessFDType = /* PROX_FDTYPE_FSEVENTS  */ 7
	ProcessFDTypeNetPolicy ProcessFDType = /* PROX_FDTYPE_NETPOLICY */ 9
)
