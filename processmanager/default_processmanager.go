package processmanager

import "errors"

type DefaultProcessManager struct {
	ProcessManager
	provider Provider
}

func (dpm *DefaultProcessManager) GetProcessIDListByName(processName string) ([]int, error) {
	value := []int{}
	if processName == "" {
		return value, errors.New("invalid process name")
	}

	return dpm.provider.OnGetProcessIDListByName(processName)
}

func (dpm *DefaultProcessManager) GetProcessNameByID(processId int) (string, error) {
	return dpm.provider.OnGetProcessNameByID(processId)
}

func (dpm *DefaultProcessManager) GetProcessPathByID(processId int) (string, error) {
	return dpm.provider.OnGetProcessPathByID(processId)
}

func (dpm *DefaultProcessManager) IsCurrentProcessRunningAsNonUser() bool {
	return dpm.provider.OnIsCurrentProcessRunningAsNonUser()
}

func (dpm *DefaultProcessManager) IsCurrentProcessRunningInDockerContainer() bool {
	return dpm.provider.OnIsCurrentProcessRunningInDockerContainer()
}
