package interop

type ProcessIDFlavor DarwinInt

/* https://github.com/apple/darwin-xnu/blob/main/bsd/sys/proc_info.h */
const (
	ProcessIDFlavorListFDs             ProcessIDFlavor = /* PROC_PIDLISTFDS             */ 1
	ProcessIDFlavorTaskAllInfo         ProcessIDFlavor = /* PROC_PIDTASKALLINFO         */ 2
	ProcessIDFlavorTbsdInfo            ProcessIDFlavor = /* PROC_PIDTBSDINFO            */ 3
	ProcessIDFlavorTaskInfo            ProcessIDFlavor = /* PROC_PIDTASKINFO            */ 4
	ProcessIDFlavorThreadInfo          ProcessIDFlavor = /* PROC_PIDTHREADINFO          */ 5
	ProcessIDFlavorListThreads         ProcessIDFlavor = /* PROC_PIDLISTTHREADS         */ 6
	ProcessIDFlavorRegionInfo          ProcessIDFlavor = /* PROC_PIDREGIONINFO          */ 7
	ProcessIDFlavorRegionPathInfo      ProcessIDFlavor = /* PROC_PIDREGIONPATHINFO      */ 8
	ProcessIDFlavorVnodePathInfo       ProcessIDFlavor = /* PROC_PIDVNODEPATHINFO       */ 9
	ProcessIDFlavorThreadPathInfo      ProcessIDFlavor = /* PROC_PIDTHREADPATHINFO      */ 10
	ProcessIDFlavorPathInfo            ProcessIDFlavor = /* PROC_PIDPATHINFO            */ 11
	ProcessIDFlavorWorkQueueInfo       ProcessIDFlavor = /* PROC_PIDWORKQUEUEINFO       */ 12
	ProcessIDFlavorTshortBsdInfo       ProcessIDFlavor = /* PROC_PIDT_SHORTBSDINFO      */ 13
	ProcessIDFlavorListFilePorts       ProcessIDFlavor = /* PROC_PIDLISTFILEPORTS       */ 14
	ProcessIDFlavorThreadID64Info      ProcessIDFlavor = /* PROC_PIDTHREADID64INFO      */ 15
	ProcessIDFlavorRusage              ProcessIDFlavor = /* PROC_PID_RUSAGE             */ 16
	ProcessIDFlavorUniqIdentifierInfo  ProcessIDFlavor = /* PROC_PIDUNIQIDENTIFIERINFO  */ 17
	ProcessIDFlavorTbsdInfoWithUniqId  ProcessIDFlavor = /* PROC_PIDT_BSDINFOWITHUNIQID */ 18
	ProcessIDFlavorArchInfo            ProcessIDFlavor = /* PROC_PIDARCHINFO            */ 19
	ProcessIDFlavorCoalitionInfo       ProcessIDFlavor = /* PROC_PIDCOALITIONINFO       */ 20
	ProcessIDFlavorNoteExit            ProcessIDFlavor = /* PROC_PIDNOTEEXIT            */ 21
	ProcessIDFlavorRegionPathInfo2     ProcessIDFlavor = /* PROC_PIDREGIONPATHINFO2     */ 22
	ProcessIDFlavorRegionPathInfo3     ProcessIDFlavor = /* PROC_PIDREGIONPATHINFO3     */ 23
	ProcessIDFlavorExitReasonInfo      ProcessIDFlavor = /* PROC_PIDEXITREASONINFO      */ 24
	ProcessIDFlavorExitReasonBasicInfo ProcessIDFlavor = /* PROC_PIDEXITREASONBASICINFO */ 25
	ProcessIDFlavorListUptrs           ProcessIDFlavor = /* PROC_PIDLISTUPTRS           */ 26
	ProcessIDFlavorListDynkQueues      ProcessIDFlavor = /* PROC_PIDLISTDYNKQUEUES      */ 27
	ProcessIDFlavorListThreadIDs       ProcessIDFlavor = /* PROC_PIDLISTTHREADIDS       */ 28
	ProcessIDFlavorVmrtFaultInfo       ProcessIDFlavor = /* PROC_PIDVMRTFAULTINFO       */ 29
	ProcessIDFlavorPlatformInfo        ProcessIDFlavor = /* PROC_PIDPLATFORMINFO        */ 30
	ProcessIDFlavorRegionPath          ProcessIDFlavor = /* PROC_PIDREGIONPATH          */ 31
	ProcessIDFlavorIpcTableInfo        ProcessIDFlavor = /* PROC_PIDIPCTABLEINFO        */ 32
)
