package desktopadapter

import (
	"strings"

	"github.com/AtelierNeon/an-core-api-golang/interop"
	"github.com/AtelierNeon/an-core-api-golang/interop/darwinlibproc"
)

func (pmp *ProcessManagerProvider) OnGetProcessIDListByName(processName string) ([]int, error) {
	result := []int{}
	processIDList, err := darwinlibproc.ListRunningProcessIDs()
	if err != nil {
		return result, err
	}

	for _, processID := range processIDList {
		pid := int(processID)
		foundProcessName, err := pmp.OnGetProcessNameByID(pid)
		if err != nil {
			continue
		}
		if foundProcessName == processName {
			result = append(result, pid)
		}
	}

	return result, nil
}

func (pmp *ProcessManagerProvider) OnGetProcessNameByID(processID int) (string, error) {
	processPath, err := pmp.OnGetProcessPathByID(processID)
	if err != nil {
		return "", err
	}

	index := strings.LastIndex(processPath, "/")
	if index < 0 {
		return processPath, nil
	}

	return processPath[index+1:], nil
}

func (pmp *ProcessManagerProvider) OnGetProcessPathByID(processID int) (string, error) {
	pid := interop.ProcessID(processID)
	return darwinlibproc.GetProcessPathByProcessID(pid)
}

func (pmp *ProcessManagerProvider) OnIsCurrentProcessRunningAsNonUser() bool {
	// TODO not implemented yet
	return false
}

func (pmp *ProcessManagerProvider) OnIsCurrentProcessRunningInDockerContainer() bool {
	return false
}
