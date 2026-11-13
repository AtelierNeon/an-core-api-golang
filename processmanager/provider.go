package processmanager

type Provider interface {
	OnGetProcessIDListByName(processName string) ([]int, error)
	OnGetProcessNameByID(processID int) (string, error)
	OnGetProcessPathByID(processID int) (string, error)
	OnIsCurrentProcessRunningAsNonUser() bool
	OnIsCurrentProcessRunningInDockerContainer() bool
}
