package processmanager

type ProcessManager interface {
	GetProcessIDListByName(processName string) ([]int, error)
	GetProcessNameByID(processID int) (string, error)
	GetProcessPathByID(processID int) (string, error)
	IsCurrentProcessRunningAsNonUser() bool
	IsCurrentProcessRunningInDockerContainer() bool
}
