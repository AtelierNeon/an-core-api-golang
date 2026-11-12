package linuxprocfs

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func GetProcessNameByID(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/stat", processID)
	buffer, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	data := string(buffer)
	nameBegin := strings.IndexRune(data, '(') + 1
	nameEnd := strings.IndexRune(data[nameBegin:], ')')
	return data[nameBegin : nameBegin+nameEnd], nil
}

func GetProcessPathByID(processID int) (string, error) {
	result, err := getProcessPathByIDFromExe(processID)
	if err != nil {
		result, err = getProcessPathByIDFromCmdline(processID)
	}

	return result, err
}

func GetProcessIDListByName(processName string) ([]int, error) {
	result, err := getProcessIDListByNameFromExe(processName)
	if err != nil {
		result, err = getProcessIDListByNameFromCmdline(processName)
	}

	return result, err
}

func getProcessIDListByNameFromCmdline(processName string) ([]int, error) {
	result := []int{}
	procEntries, err := os.ReadDir("/proc")
	if err != nil {
		return result, err
	}

	for _, procEntry := range procEntries {
		processID, err := strconv.Atoi(procEntry.Name())
		if err != nil {
			continue
		}

		processPath, err := getProcessPathByIDFromCmdline(processID)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return result, err
		}

		if strings.HasSuffix(processPath, "/"+processName) {
			result = append(result, processID)
		}
	}

	return result, nil
}

func getProcessIDListByNameFromExe(processName string) ([]int, error) {
	result := []int{}
	procEntries, err := os.ReadDir("/proc")
	if err != nil {
		return result, err
	}

	for _, procEntry := range procEntries {
		processID, err := strconv.Atoi(procEntry.Name())
		if err != nil {
			continue
		}

		processPath, err := getProcessPathByIDFromExe(processID)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return result, err
		}

		if strings.HasSuffix(processPath, "/"+processName) {
			result = append(result, processID)
		}
	}

	return result, nil
}

func getProcessPathByIDFromCmdline(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/cmdline", processID)

	file, err := os.OpenFile(path, os.O_RDONLY, 0o440)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := io.LimitReader(file, 1024)
	buffer, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	data := strings.Split(string(buffer), "\x00")
	return strings.TrimSpace(data[0]), nil
}

func getProcessPathByIDFromExe(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/exe", processID)

	var buffer [1024]byte
	bufferRead, err := unix.Readlink(path, buffer[:])
	if err != nil {
		return "", err
	}

	return string(buffer[:bufferRead]), nil
}
