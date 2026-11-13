package interop

type ProcessInfoType uint32

/* https://github.com/apple/darwin-xnu/blob/main/bsd/sys/proc_info.h */
const (
	ProcessInfoTypeAllPIDs  ProcessInfoType = /* PROC_ALL_PIDS  */ 1
	ProcessInfoTypePGRPOnly ProcessInfoType = /* PROC_PGRP_ONLY */ 2
	ProcessInfoTypeTTYOnly  ProcessInfoType = /* PROC_TTY_ONLY  */ 3
	ProcessInfoTypeUIDOnly  ProcessInfoType = /* PROC_UID_ONLY  */ 4
	ProcessInfoTypeRUIDOnly ProcessInfoType = /* PROC_RUID_ONLY */ 5
	ProcessInfoTypePPIDOnly ProcessInfoType = /* PROC_PPID_ONLY */ 6
	ProcessInfoTypeKDBGOnly ProcessInfoType = /* PROC_KDBG_ONLY */ 7
)
