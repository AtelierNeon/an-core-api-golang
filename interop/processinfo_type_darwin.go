package interop

type ProcessInfoType uint32

/* https://github.com/apple/darwin-xnu/blob/main/bsd/sys/proc_info.h */
const (
	processInfoTypeAllPIDs  ProcessInfoType = /* PROC_ALL_PIDS  */ 1
	processInfoTypePGRPOnly ProcessInfoType = /* PROC_PGRP_ONLY */ 2
	processInfoTypeTTYOnly  ProcessInfoType = /* PROC_TTY_ONLY  */ 3
	processInfoTypeUIDOnly  ProcessInfoType = /* PROC_UID_ONLY  */ 4
	processInfoTypeRUIDOnly ProcessInfoType = /* PROC_RUID_ONLY */ 5
	processInfoTypePPIDOnly ProcessInfoType = /* PROC_PPID_ONLY */ 6
	processInfoTypeKDBGOnly ProcessInfoType = /* PROC_KDBG_ONLY */ 7
)
