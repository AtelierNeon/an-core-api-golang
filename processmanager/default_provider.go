package processmanager

import "errors"

type DefaultProvider struct {
	Provider
}

func (dp *DefaultProvider) OnGetProcessIDListByName(processName string) ([]int, error) {
	return nil, errors.New("not implemented yet")
}

func (dp *DefaultProvider) OnGetProcessNameByID(processID int) (string, error) {
	return "", errors.New("not implemented yet")
}

func (dp *DefaultProvider) OnGetProcessPathByID(processID int) (string, error) {
	return "", errors.New("not implemented yet")
}

func (dp *DefaultProvider) OnIsCurrentProcessRunningAsNonUser() bool {
	return false
}

func (dp *DefaultProvider) OnIsCurrentProcessRunningInDockerContainer() bool {
	return false
}
