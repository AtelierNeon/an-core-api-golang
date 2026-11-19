package linuxprocfs

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AtelierNeon/an-core-api-golang/interop"
)

func decodeIPv4AddressAndPort(rawAddressPort string) (string, uint16, error) {
	tokens := strings.Split(rawAddressPort, ":")
	if len(tokens) < 2 {
		return "", 0, fmt.Errorf("raw address does not contain port, %s", rawAddressPort)
	}

	rawPort := tokens[1]
	port, err := strconv.ParseUint(rawPort, 16, 16)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port format: %s", rawPort)
	}

	rawAddress := tokens[0]
	reversedAddressInBytes, err := hex.DecodeString(rawAddress)
	if err != nil {
		return "", 0, fmt.Errorf("decode raw address: %s in hex error", rawAddress)
	}

	addressInBytes := reverseInBytes(reversedAddressInBytes)
	ip := net.IP(addressInBytes)
	return ip.String(), uint16(port), nil
}

func decodeIPv6AddressAndPort(rawAddressPort string) (string, uint16, error) {
	tokens := strings.Split(rawAddressPort, ":")
	if len(tokens) < 2 {
		return "", 0, fmt.Errorf("raw address does not contain port, %s", rawAddressPort)
	}

	rawPort := tokens[1]
	port, err := strconv.ParseUint(rawPort, 16, 16)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port format: %s", rawAddressPort)
	}

	rawAddress := tokens[0]
	reversedAddressInBytes, err := hex.DecodeString(rawAddress)
	if err != nil {
		return rawAddress, 0, fmt.Errorf("decode raw address: %s in hex error", rawAddressPort)
	}

	if len(reversedAddressInBytes) != 16 {
		return "", 0, fmt.Errorf("invalid ipv6 string, %s", rawAddressPort)
	}

	addressInBytes := make([]byte, 0, 16)
	for i := 0; i < len(reversedAddressInBytes); i += 4 {
		reversedSegmentInBytes := reversedAddressInBytes[i : i+4]
		segmentInBytes := reverseInBytes(reversedSegmentInBytes)
		addressInBytes = append(addressInBytes, segmentInBytes...)
	}
	ip := net.IP(addressInBytes)
	return ip.String(), uint16(port), nil
}

func decodeTcpState(rawState string) (uint32, error) {
	state, stateErr := strconv.ParseUint(rawState, 16, 16)
	if stateErr != nil {
		return 0, fmt.Errorf("invalid state: %s", rawState)
	}

	return uint32(state), nil
}

func GetPortBindingList() ([]PortBindingInfo, error) {
	processIDList, err := getRunningProcessIDList()
	if err != nil {
		return nil, err
	}

	inodeMap := make(map[string]int)
	for _, processID := range processIDList {
		inodeMapForProcessID, err := getSocketInodeToProcessIdMapForProcessID(processID)
		if err != nil {
			continue
		}

		for inode, inodeProcessID := range inodeMapForProcessID {
			inodeMap[inode] = inodeProcessID
		}
	}

	result := []PortBindingInfo{}

	tcp4List, err := getPortBindingListForTCP4()
	if err != nil {
		return nil, err
	}
	result = append(result, tcp4List...)

	tcp6List, err := getPortBindingListForTCP6()
	if err != nil {
		return nil, err
	}
	result = append(result, tcp6List...)

	udp4List, err := getPortBindingListForUDP4()
	if err != nil {
		return nil, err
	}
	result = append(result, udp4List...)

	udp6List, err := getPortBindingListForUDP6()
	if err != nil {
		return nil, err
	}
	result = append(result, udp6List...)

	for i, item := range result {
		inode := item.Inode
		processID, exists := inodeMap[inode]
		if exists {
			result[i].OwnerProcessID = processID
		}
	}

	return result, nil
}

func getPortBindingListForTCP4() ([]PortBindingInfo, error) {
	/*
	  Example content for /proc/self/net/tcp

	  sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
	   0: 0100007F:0277 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 445336 1 0000000000000000 100 0 0 10 0
	   1: 00000000:0016 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 7452 1 0000000000000000 100 0 0 10 0
	   2: 3500007F:0035 00000000:0000 0A 00000000:00000000 00:00000000 00000000   101        0 5952 1 0000000000000000 100 0 0 10 5
	   3: 0100007F:B347 00000000:0000 0A 00000000:00000000 00:00000000 00000000  1000        0 459972 1 0000000000000000 100 0 0 10 0
	   4: 50884D64:8436 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 11331 1 0000000000000000 100 0 0 10 0
	   5: 50884D64:0016 6BD05564:FA80 01 00000160:00000000 01:0000000C 00000000     0        0 457589 4 0000000000000000 26 4 31 10 -1
	   6: 50884D64:0016 6BD05564:FA90 01 00000000:00000000 02:00091D1F 00000000     0        0 458925 2 0000000000000000 28 4 30 10 115
	   7: 0100007F:D974 0100007F:B347 01 00000000:00000000 00:00000000 00000000  1000        0 459979 1 0000000000000000 20 4 30 10 -1
	   8: 0100007F:B347 0100007F:D974 01 00000000:00000000 00:00000000 00000000  1000        0 459980 1 0000000000000000 20 4 30 10 -1
	   9: 31A8A8C0:E77E 1E6993CD:01BB 01 00000000:00000000 02:000006DA 00000000     0        0 459431 2 0000000000000000 25 4 30 10 -1
	  10: 31A8A8C0:A7D6 39EF4903:01BB 01 00000000:00000000 02:00000475 00000000     0        0 451966 2 0000000000000000 46 4 30 10 -1
	  11: 31A8A8C0:8C7C BC9B5467:01BB 01 00000000:00000000 02:00000365 00000000     0        0 453977 2 0000000000000000 26 5 30 10 -1
	  12: 31A8A8C0:BE04 BBCACF34:01BB 01 00000000:00000000 02:000006F6 00000000     0        0 455664 2 0000000000000000 41 4 30 10 -1
	*/
	content, err := os.ReadFile("/proc/self/net/tcp")
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines[1:] {
		field := strings.Fields(string(line))
		if len(field) < 10 {
			continue
		}

		localAddress, localPort, err := decodeIPv4AddressAndPort(field[1])
		if err != nil {
			continue
		}

		remoteAddress, remotePort, err := decodeIPv4AddressAndPort(field[2])
		if err != nil {
			continue
		}

		tcpState, err := decodeTcpState(field[3])
		if err != nil {
			continue
		}

		item := PortBindingInfo{
			LocalAddress:  localAddress,
			LocalPort:     localPort,
			RemoteAddress: remoteAddress,
			RemotePort:    remotePort,
			State:         interop.ProcfsTcpState(tcpState).String(),
			Kind:          "tcp4",
			Inode:         field[9],
		}
		result = append(result, item)
	}

	return result, nil
}

func getPortBindingListForTCP6() ([]PortBindingInfo, error) {
	/*
	  Example content for /proc/self/net/tcp6

	  sl  local_address                         remote_address                        st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
	   0: 5C117AFD0000E0A100000000518801D6:C004 00000000000000000000000000000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 11332 1 0000000000000000 100 0 0 10 0
	   1: 00000000000000000000000000000000:0016 00000000000000000000000000000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 7454 1 0000000000000000 100 0 0 10 0
	   2: 00000000000000000000000000000000:0D3D 00000000000000000000000000000000:0000 0A 00000000:00000000 00:00000000 00000000  1000        0 15729 1 0000000000000000 100 0 0 10 0
	   3: 00000000000000000000000001000000:0277 00000000000000000000000000000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 445335 1 0000000000000000 100 0 0 10 0
	*/
	content, err := os.ReadFile("/proc/self/net/tcp6")
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines[1:] {
		field := strings.Fields(string(line))
		if len(field) < 10 {
			continue
		}

		localAddress, localPort, err := decodeIPv6AddressAndPort(field[1])
		if err != nil {
			continue
		}

		remoteAddress, remotePort, err := decodeIPv6AddressAndPort(field[2])
		if err != nil {
			continue
		}

		tcpState, err := decodeTcpState(field[3])
		if err != nil {
			continue
		}

		item := PortBindingInfo{
			LocalAddress:  localAddress,
			LocalPort:     localPort,
			RemoteAddress: remoteAddress,
			RemotePort:    remotePort,
			State:         interop.ProcfsTcpState(tcpState).String(),
			Kind:          "tcp6",
			Inode:         field[9],
		}
		result = append(result, item)
	}

	return result, nil
}

func getPortBindingListForUDP4() ([]PortBindingInfo, error) {
	/*
	  Example content for /proc/self/net/udp

	    sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode ref pointer drops
	  2154: 3500007F:0035 00000000:0000 07 00000000:00000000 00:00000000 00000000   101        0 5951 2 0000000000000000 0
	  2169: 31A8A8C0:0044 01A8A8C0:0043 01 00000000:00000000 00:00000000 00000000     0        0 4956 2 0000000000000000 0
	  2782: 00000000:A2A9 00000000:0000 07 00000000:00000000 00:00000000 00000000     0        0 91764 2 0000000000000000 0
	  5893: 00000000:CED0 00000000:0000 07 00000000:00000000 00:00000000 00000000   114        0 6740 2 0000000000000000 0
	  7454: 00000000:14E9 00000000:0000 07 00000000:00000000 00:00000000 00000000   114        0 6738 2 0000000000000000 0
	*/
	content, err := os.ReadFile("/proc/self/net/udp")
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines[1:] {
		field := strings.Fields(string(line))
		if len(field) < 10 {
			continue
		}

		localAddress, localPort, err := decodeIPv4AddressAndPort(field[1])
		if err != nil {
			continue
		}

		remoteAddress, remotePort, err := decodeIPv4AddressAndPort(field[2])
		if err != nil {
			continue
		}

		item := PortBindingInfo{
			LocalAddress:  localAddress,
			LocalPort:     localPort,
			RemoteAddress: remoteAddress,
			RemotePort:    remotePort,
			State:         "",
			Kind:          "udp4",
			Inode:         field[9],
		}
		result = append(result, item)
	}

	return result, nil
}

func getPortBindingListForUDP6() ([]PortBindingInfo, error) {
	/*
	  Example content for /proc/self/net/udp6

	    sl  local_address                         remote_address                        st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode ref pointer drops
	  2782: 00000000000000000000000000000000:A2A9 00000000000000000000000000000000:0000 07 00000000:00000000 00:00000000 00000000     0        0 91763 2 0000000000000000 0
	  6480: 00000000000000000000000000000000:911B 00000000000000000000000000000000:0000 07 00000000:00000000 00:00000000 00000000   114        0 6741 2 0000000000000000 0
	  7454: 00000000000000000000000000000000:14E9 00000000000000000000000000000000:0000 07 00000000:00000000 00:00000000 00000000   114        0 6739 2 0000000000000000 0
	*/
	content, err := os.ReadFile("/proc/self/net/udp6")
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines[1:] {
		field := strings.Fields(string(line))
		if len(field) < 10 {
			continue
		}

		localAddress, localPort, err := decodeIPv6AddressAndPort(field[1])
		if err != nil {
			continue
		}

		remoteAddress, remotePort, err := decodeIPv6AddressAndPort(field[2])
		if err != nil {
			continue
		}

		item := PortBindingInfo{
			LocalAddress:  localAddress,
			LocalPort:     localPort,
			RemoteAddress: remoteAddress,
			RemotePort:    remotePort,
			State:         "",
			Kind:          "udp6",
			Inode:         field[9],
		}
		result = append(result, item)
	}

	return result, nil
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
	procEntries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	result := []int{}

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

			return nil, err
		}

		if strings.HasSuffix(processPath, "/"+processName) {
			result = append(result, processID)
		}
	}

	return result, nil
}

func GetProcessNameByID(processID int) (string, error) {
	processName, err := getProcessNameByIDFromExe(processID)
	if err != nil {
		processName, err = getProcessNameByIDFromCmdline(processID)
	}
	if err != nil {
		processName, err = getProcessNameByIDFromStat(processID)
	}

	if err != nil {
		return "", err
	}
	return processName, nil
}

func getProcessNameByIDFromCmdline(processID int) (string, error) {
	path, err := getProcessPathByIDFromCmdline(processID)
	if err != nil {
		return "", err
	}

	return filepath.Base(path), nil
}

func getProcessNameByIDFromExe(processID int) (string, error) {
	path, err := getProcessPathByIDFromExe(processID)
	if err != nil {
		return "", err
	}

	return filepath.Base(path), nil
}

func getProcessNameByIDFromStat(processID int) (string, error) {
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

func getProcessPathByIDFromCmdline(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/cmdline", processID)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	tokens := strings.Split(string(data), "\x00")
	result := strings.TrimSpace(tokens[0])
	if result == "" {
		return "", os.ErrNotExist
	}
	return result, nil
}

func getProcessPathByIDFromExe(processID int) (string, error) {
	path := fmt.Sprintf("/proc/%d/exe", processID)
	return os.Readlink(path)
}

func getRunningProcessIDList() ([]int, error) {
	dirEntries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	result := []int{}

	for _, dirEntry := range dirEntries {
		pid, err := strconv.ParseInt(dirEntry.Name(), 10, 32)
		if err != nil {
			continue
		}

		result = append(result, int(pid))
	}

	return result, nil
}

func getSocketInodeToProcessIdMapForProcessID(processID int) (map[string]int, error) {
	path := fmt.Sprintf("/proc/%d/fd", processID)
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make(map[string]int)

	for _, dirEntry := range dirEntries {
		inodePath := fmt.Sprintf("/proc/%d/fd/%s", processID, dirEntry.Name())
		/*
		  Example inode

		  lr-x------ 1 install install 64 Mar 19 10:09 0 -> /dev/null
		  l-wx------ 1 install install 64 Mar 19 10:09 1 -> /home/install/.vscode-server/.cli.ddc367ed5c8936efe395cffeec279b04ffd7db78.log
		  l-wx------ 1 install install 64 Mar 19 10:09 2 -> 'anon_inode:[eventpoll]'
		  lrwx------ 1 install install 64 Mar 19 10:09 3 -> 'anon_inode:[eventfd]'
		  lrwx------ 1 install install 64 Mar 19 10:09 4 -> 'anon_inode:[eventpoll]'
		  lrwx------ 1 install install 64 Mar 19 10:09 5 -> 'socket:[459963]'
		  lrwx------ 1 install install 64 Mar 19 10:09 6 -> 'socket:[459964]'
		*/
		inode, err := os.Readlink(inodePath)
		if err != nil {
			continue
		}
		if !strings.HasPrefix(inode, "socket:[") {
			continue
		}

		l := len(inode)
		inode = inode[8 : l-1]

		result[inode] = processID
	}
	return result, nil
}

func reverseInBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}
