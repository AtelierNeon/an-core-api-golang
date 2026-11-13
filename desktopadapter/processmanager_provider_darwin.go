package desktopadapter

import (
	"strings"

	"github.com/AtelierNeon/an-core-api-golang/interop"
	"github.com/AtelierNeon/an-core-api-golang/interop/darwinlibproc"
)

func (pmp *ProcessManagerProvider) OnGetProcessIDListByName(processName string) ([]int, error) {
	value := []int{}
	processIdList, processIdListErr := darwinlibproc.GetRunningProcessIDList()
	if processIdListErr != nil {
		return value, processIdListErr
	}

	for _, processId := range processIdList {
		pid := int(processId)
		processName2, processName2Err := pmp.OnGetProcessNameByID(pid)
		if processName2Err != nil {
			continue
		}
		if processName2 == processName {
			value = append(value, pid)
		}
	}

	return value, nil
}

func (pmp *ProcessManagerProvider) OnGetProcessNameByID(processID int) (string, error) {
	processPath, processPathErr := pmp.OnGetProcessPathByID(processID)
	if processPathErr != nil {
		return "", processPathErr
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
