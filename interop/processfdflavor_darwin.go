package interop

/* https://github.com/apple-oss-distributions/xnu/blob/main/libsyscall/wrappers/libproc/libproc.h */
type ProcessFDFlavor DarwinInt

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
const (
	ProcessFDFlavorVnode     ProcessFDFlavor = /* PROC_PIDFDVNODEINFO      */ 1
	ProcessFDFlavorVnodePath ProcessFDFlavor = /* PROC_PIDFDVNODEPATHINFO  */ 2
	ProcessFDFlavorSocket    ProcessFDFlavor = /* PROC_PIDFDSOCKETINFO     */ 3
	ProcessFDFlavorPsem      ProcessFDFlavor = /* PROC_PIDFDPSEMINFO       */ 4
	ProcessFDFlavorPshm      ProcessFDFlavor = /* PROC_PIDFDPSHMINFO       */ 5
	ProcessFDFlavorPipe      ProcessFDFlavor = /* PROC_PIDFDPIPEINFO       */ 6
	ProcessFDFlavorKqueue    ProcessFDFlavor = /* PROC_PIDFDKQUEUEINFO     */ 7
	ProcessFDFlavorAtalk     ProcessFDFlavor = /* PROC_PIDFDATALKINFO      */ 8
	ProcessFDFlavorKqueueExt ProcessFDFlavor = /* PROC_PIDFDKQUEUE_EXTINFO */ 9
)
