package interop

type ProcessIDListType uint32

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
const (
	ProcessIDListTypeAllPIDs  ProcessIDListType = /* PROC_ALL_PIDS  */ 1
	ProcessIDListTypePGRPOnly ProcessIDListType = /* PROC_PGRP_ONLY */ 2
	ProcessIDListTypeTTYOnly  ProcessIDListType = /* PROC_TTY_ONLY  */ 3
	ProcessIDListTypeUIDOnly  ProcessIDListType = /* PROC_UID_ONLY  */ 4
	ProcessIDListTypeRUIDOnly ProcessIDListType = /* PROC_RUID_ONLY */ 5
	ProcessIDListTypePPIDOnly ProcessIDListType = /* PROC_PPID_ONLY */ 6
	ProcessIDListTypeKDBGOnly ProcessIDListType = /* PROC_KDBG_ONLY */ 7
)
