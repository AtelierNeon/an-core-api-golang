package desktopadapter

import (
	"os"

	"github.com/AtelierNeon/an-core-api-golang/interop"
	"github.com/AtelierNeon/an-core-api-golang/interop/linuxprocfs"
)

func (pmp *ProcessManagerProvider) OnGetProcessIDListByName(processName string) ([]int, error) {
	return linuxprocfs.GetProcessIDListByName(processName)
}

func (pmp *ProcessManagerProvider) OnGetProcessNameByID(processID int) (string, error) {
	return linuxprocfs.GetProcessNameByID(processID)
}

func (pmp *ProcessManagerProvider) OnGetProcessPathByID(processID int) (string, error) {
	return linuxprocfs.GetProcessPathByID(processID)
}

func (pmp *ProcessManagerProvider) OnIsCurrentProcessRunningAsNonUser() bool {
	if pmp.OnIsCurrentProcessRunningInDockerContainer() {
		return true
	}

	initProcessID := interop.GetInitProcessID()
	parentProcessID := os.Getppid()
	if parentProcessID == initProcessID {
		return true
	}

	parentProcessName, err := pmp.OnGetProcessNameByID(parentProcessID)
	if err != nil {
		return false
	}

	return parentProcessName == "systemd"
}

func (pmp *ProcessManagerProvider) OnIsCurrentProcessRunningInDockerContainer() bool {
	_, err := os.Stat("/.dockerenv")
	return err == nil
}
