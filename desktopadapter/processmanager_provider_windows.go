package desktopadapter

import (
	"fmt"
	"unsafe"

	"github.com/AtelierNeon/an-core-api-golang/interop"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc"
)

func (pmp *ProcessManagerProvider) OnGetProcessIDListByName(processName string) ([]int, error) {
	result := []int{}
	handle, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return result, err
	}
	defer windows.CloseHandle(handle)

	var processEntry windows.ProcessEntry32
	processEntry.Size = uint32(unsafe.Sizeof(processEntry))
	err = windows.Process32First(handle, &processEntry)
	if err != nil {
		return result, err
	}

	if windows.UTF16ToString(processEntry.ExeFile[:]) == processName {
		result = append(result, int(processEntry.ProcessID))
	}
	for {
		err = windows.Process32Next(handle, &processEntry)
		if err == windows.ERROR_NO_MORE_FILES {
			break
		}
		if err != nil {
			return result, err
		}

		if windows.UTF16ToString(processEntry.ExeFile[:]) == processName {
			result = append(result, int(processEntry.ProcessID))
		}
	}

	return result, nil
}

func (pmp *ProcessManagerProvider) OnGetProcessNameByID(processID int) (string, error) {
	handle, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return "", err
	}
	defer windows.CloseHandle(handle)

	var processEntry windows.ProcessEntry32
	processEntry.Size = uint32(unsafe.Sizeof(processEntry))
	err = windows.Process32First(handle, &processEntry)
	if err != nil {
		return "", err
	}

	if processEntry.ProcessID == uint32(processID) {
		return windows.UTF16ToString(processEntry.ExeFile[:]), nil
	}

	for {
		err = windows.Process32Next(handle, &processEntry)
		if err != nil {
			return "", err
		}

		if processEntry.ProcessID == uint32(processID) {
			return windows.UTF16ToString(processEntry.ExeFile[:]), nil
		}
	}
}

func (pmp *ProcessManagerProvider) OnGetProcessPathByID(processID int) (string, error) {
	handle, err := windows.OpenProcess(
		windows.PROCESS_QUERY_LIMITED_INFORMATION,
		false,
		uint32(processID),
	)
	defer windows.CloseHandle(handle)

	if err != nil {
		return "", fmt.Errorf("can not open process: %v", err)
	}

	initPathSize := 256
	maxPathSize := interop.NTFS_MAX_PATH
	for capacity := initPathSize; capacity <= maxPathSize; capacity *= 2 {
		buffer := make([]uint16, capacity)
		bufferSize := uint32(len(buffer))
		err = windows.QueryFullProcessImageName(
			handle,
			0,
			&buffer[0],
			&bufferSize,
		)

		if err == nil {
			return windows.UTF16ToString(buffer[:bufferSize]), nil
		}

		if err != windows.ERROR_INSUFFICIENT_BUFFER || capacity >= maxPathSize {
			return "", fmt.Errorf("can not query full process image name: %v", err)
		}
	}

	return "", fmt.Errorf("can not query full process image name: %v", err)
}

func (pmp *ProcessManagerProvider) OnIsCurrentProcessRunningAsNonUser() bool {
	isService, err := svc.IsWindowsService()
	if err != nil {
		return false
	}

	return isService
}

func (pmp *ProcessManagerProvider) OnIsCurrentProcessRunningInDockerContainer() bool {
	return false
}
