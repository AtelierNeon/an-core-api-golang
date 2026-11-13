package linuxprocfs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetProcessNameByID(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/stat", processID)
	buffer, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	data := string(buffer)
	start := strings.IndexByte(data, '(')
	end := strings.LastIndexByte(data, ')')
	if start < 0 || end <= start {
		return "", fmt.Errorf("invalid process stat format: %s", path)
	}

	return data[start+1 : end], nil
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
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	tokens := strings.Split(string(data), "\x00")
	return strings.TrimSpace(tokens[0]), nil
}

func getProcessPathByIDFromExe(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/exe", processID)
	return os.Readlink(path)
}
