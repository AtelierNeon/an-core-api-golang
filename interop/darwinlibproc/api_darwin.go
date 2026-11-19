package darwinlibproc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"unsafe"

	"github.com/AtelierNeon/an-core-api-golang/interop"
)

const (
	procPidPathInfoSize = 0x400
)

func decodeIPv4Address(data *byte) string {
	in4In6Addr := (*interop.In4In6_addr)(unsafe.Pointer(data))
	ip := net.IP(in4In6Addr.Addr4[:])
	return ip.String()
}

func decodeIPv6Address(data *byte) string {
	in6Addr := *(*[]byte)(unsafe.Pointer(data))
	ip := net.IP(in6Addr)
	return ip.String()
}

func decodeInSocketInfo(data *byte) *interop.In_sockinfo {
	return (*interop.In_sockinfo)(unsafe.Pointer(data))
}

func decodeTcpSocketInfo(data *byte) *interop.Tcp_sockinfo {
	return (*interop.Tcp_sockinfo)(unsafe.Pointer(data))
}

func GetPortBindingList() ([]PortBindingInfo, error) {
	result := []PortBindingInfo{}

	processIDList, err := ListRunningProcessIDs()
	if err != nil {
		return result, err
	}

	for _, processID := range processIDList {
		processFdMap, err := getProcessFdMapByProcessID(processID)
		if err != nil {
			continue
		}

		for _, processFdInfo := range processFdMap {
			processFdType := processFdInfo.ProcFdType
			if processFdType != interop.ProcessFdTypeSocket {
				continue
			}

			socketFdInfo, err := getSocketFdInfo(
				processID,
				processFdInfo,
			)
			if err != nil {
				continue
			}

			family := socketFdInfo.Socket.Family
			kind := socketFdInfo.Socket.Kind
			if family == interop.AddressFamilyInet {
				if kind == interop.SocketKindTcp {
					tcpSocketInfo := decodeTcpSocketInfo(&socketFdInfo.Socket.Proto[0])
					inSockInfo := tcpSocketInfo.Ini
					item := PortBindingInfo{
						LocalAddress:   decodeIPv4Address(&inSockInfo.Laddr[0]),
						LocalPort:      netToNativeEndianInShort(uint16(inSockInfo.Lport)),
						RemoteAddress:  decodeIPv4Address(&inSockInfo.Faddr[0]),
						RemotePort:     netToNativeEndianInShort(uint16(inSockInfo.Fport)),
						OwnerProcessID: processID,
						State:          tcpSocketInfo.State.String(),
						Kind:           "tcp4",
					}
					result = append(result, item)
				} else if kind == interop.SocketKindIn {
					inSockInfo := decodeInSocketInfo(&socketFdInfo.Socket.Proto[0])
					item := PortBindingInfo{
						LocalAddress:   decodeIPv4Address(&inSockInfo.Laddr[0]),
						LocalPort:      netToNativeEndianInShort(uint16(inSockInfo.Lport)),
						RemoteAddress:  decodeIPv4Address(&inSockInfo.Faddr[0]),
						RemotePort:     netToNativeEndianInShort(uint16(inSockInfo.Fport)),
						OwnerProcessID: processID,
						State:          interop.TcpStateListen.String(),
						Kind:           "udp4",
					}
					result = append(result, item)
				}
			} else if family == interop.AddressFamilyInet6 {
				if kind == interop.SocketKindTcp {
					tcpSocketInfo := decodeTcpSocketInfo(&socketFdInfo.Socket.Proto[0])
					inSockInfo := tcpSocketInfo.Ini
					item := PortBindingInfo{
						LocalAddress:   decodeIPv6Address(&inSockInfo.Laddr[0]),
						LocalPort:      netToNativeEndianInShort(uint16(inSockInfo.Lport)),
						RemoteAddress:  decodeIPv6Address(&inSockInfo.Faddr[0]),
						RemotePort:     netToNativeEndianInShort(uint16(inSockInfo.Fport)),
						OwnerProcessID: processID,
						State:          tcpSocketInfo.State.String(),
						Kind:           "tcp6",
					}
					result = append(result, item)
				} else if kind == interop.SocketKindIn {
					inSockInfo := decodeInSocketInfo(&socketFdInfo.Socket.Proto[0])
					item := PortBindingInfo{
						LocalAddress:   decodeIPv6Address(&inSockInfo.Laddr[0]),
						LocalPort:      netToNativeEndianInShort(uint16(inSockInfo.Lport)),
						RemoteAddress:  decodeIPv6Address(&inSockInfo.Faddr[0]),
						RemotePort:     netToNativeEndianInShort(uint16(inSockInfo.Fport)),
						OwnerProcessID: processID,
						State:          interop.TcpStateListen.String(),
						Kind:           "udp6",
					}
					result = append(result, item)
				}
			}
		}
	}
	return result, nil
}

func getProcessFdMapByProcessID(processID interop.ProcessID) (map[string]interop.Proc_fdinfo, error) {
	r1, _, errno := interop.Proc_pidinfo(
		processID,
		interop.ProcessIDFlavorListFDs,
		0,
		nil,
		0,
	)
	kernelReturn := interop.KernelReturn(errno)
	if kernelReturn != interop.KernelReturnSuccess {
		return nil, errors.New(kernelReturn.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to get process info size")
	}

	bufferSize := uint32(r1)
	if bufferSize < math.MaxUint32/2 {
		bufferSize *= 2
	}
	buffer := make([]byte, bufferSize)
	r1, _, errno = interop.Proc_pidinfo(
		processID,
		interop.ProcessIDFlavorListFDs,
		0,
		unsafe.Pointer(&buffer[0]),
		bufferSize,
	)
	kernelReturn = interop.KernelReturn(errno)
	if kernelReturn != interop.KernelReturnSuccess {
		return nil, errors.New(kernelReturn.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to get process fd map")
	}

	bufferSize = uint32(r1)
	var processFdInfo interop.Proc_fdinfo
	processFdInfoListSize := int(bufferSize) / binary.Size(processFdInfo)
	result := make(map[string]interop.Proc_fdinfo)
	binaryBuffer := bytes.NewBuffer(buffer)

	for i := 0; i < processFdInfoListSize; i++ {
		if err := binary.Read(binaryBuffer, binary.NativeEndian, &processFdInfo); err != nil {
			return nil, err
		}

		key := fmt.Sprintf("%d_%d", processFdInfo.ProcFd, processFdInfo.ProcFdType)
		result[key] = processFdInfo
	}
	return result, nil
}

func getSocketFdInfo(processId interop.ProcessID, processFdInfo interop.Proc_fdinfo) (*interop.Socket_fdinfo, error) {
	bufferSize := uint32(unsafe.Sizeof(interop.Socket_fdinfo{}))
	buffer := make([]byte, bufferSize)
	r1, _, errno := interop.Proc_pidfdinfo(
		processId,
		processFdInfo.ProcFd,
		interop.ProcessFDFlavorSocket,
		unsafe.Pointer(&buffer[0]),
		bufferSize,
	)
	kernelReturn := interop.KernelReturn(errno)
	if kernelReturn != interop.KernelReturnSuccess {
		return nil, errors.New(kernelReturn.String())
	}

	if r1 == 0 {
		return nil, errors.New("invalid socketFdInfo")
	}

	result := (*interop.Socket_fdinfo)(unsafe.Pointer(&buffer[0]))
	return result, nil
}

func GetProcessPathByProcessID(processID interop.ProcessID) (string, error) {
	buffer := newProcPidpathBuffer()
	r1, _, errno := interop.Proc_pidpath(
		processID,
		unsafe.Pointer(&buffer[0]),
		uint32(len(buffer)),
	)
	kernelReturn := interop.KernelReturn(errno)
	if kernelReturn != interop.KernelReturnSuccess {
		return "", errors.New(kernelReturn.String())
	}

	if r1 == 0 {
		return "", errors.New("invalid process path size")
	}

	return interop.BufferToString(buffer), nil
}

func ListRunningProcessIDs() ([]interop.ProcessID, error) {
	r1, _, errno := interop.Proc_listpids(
		interop.ProcessIDListTypeAllPIDs,
		0,
		nil,
		0,
	)
	kernelReturn := interop.KernelReturn(errno)
	if kernelReturn != interop.KernelReturnSuccess {
		return nil, errors.New(kernelReturn.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to list running process IDs")
	}

	bufferSize := uint32(r1)
	if bufferSize < math.MaxUint32/2 {
		bufferSize *= 2
	}
	buffer := make([]byte, bufferSize)
	r1, _, errno = interop.Proc_listpids(
		interop.ProcessIDListTypeAllPIDs,
		0,
		unsafe.Pointer(&buffer[0]),
		bufferSize,
	)
	kernelReturn = interop.KernelReturn(errno)
	if kernelReturn != interop.KernelReturnSuccess {
		return nil, errors.New(kernelReturn.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to list running process IDs")
	}

	bufferSize = uint32(r1)
	var processID interop.ProcessID
	processIDListSize := int(bufferSize) / binary.Size(processID)
	result := make([]interop.ProcessID, 0, processIDListSize)
	binaryBuffer := bytes.NewBuffer(buffer)

	for i := 0; i < processIDListSize; i++ {
		if err := binary.Read(binaryBuffer, binary.NativeEndian, &processID); err != nil {
			return nil, err
		}
		if processID == 0 {
			continue
		}

		result = append(result, processID)
	}
	return result, nil
}

func netToNativeEndianInShort(data uint16) uint16 {
	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, data)
	return binary.NativeEndian.Uint16(buffer)
}

func newProcPidpathBuffer() []byte {
	return make([]byte, procPidPathInfoSize)
}
