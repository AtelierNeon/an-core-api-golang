package windowsmib

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"

	"github.com/AtelierNeon/an-core-api-golang/interop"
	"golang.org/x/sys/windows"
)

func decodeIPv4Address(addr uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", addr&255, addr>>8&255, addr>>16&255, addr>>24&255)
}

func decodeIPv6Address(addr [16]byte) string {
	var buffer [16]byte
	for i := 0; i < 16; i++ {
		buffer[i] = uint8(addr[i])
	}

	ip := net.IP(buffer[:])
	return ip.String()
}

func decodePort(port uint32) uint16 {
	return syscall.Ntohs(uint16(port))
}

func GetPortBindingList() ([]PortBindingInfo, error) {
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

	return result, nil
}

func getPortBindingListForTCP4() ([]PortBindingInfo, error) {
	var table *interop.MibTcpTableOwnerPid
	tableBufferSize := uint32(0)

	err := interop.GetExtendedTcpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET,
		interop.TcpTableClassOwnerPidAll,
		0,
	)
	if err != nil && err != windows.ERROR_INSUFFICIENT_BUFFER {
		return nil, err
	}

	tableBuffer := make([]byte, tableBufferSize)
	table = (*interop.MibTcpTableOwnerPid)(unsafe.Pointer(&tableBuffer[0]))
	err = interop.GetExtendedTcpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET,
		interop.TcpTableClassOwnerPidAll,
		0,
	)
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	if table.DwNumEntries == 0 {
		return result, nil
	}

	tableRowIndex := int(unsafe.Sizeof(table.DwNumEntries))
	tableRowStep := int(unsafe.Sizeof(table.Table))
	for i := 0; i < int(table.DwNumEntries); i++ {
		tableRow := (*interop.MibTcpRowOwnerPid)(unsafe.Pointer(&tableBuffer[tableRowIndex]))
		item := PortBindingInfo{
			LocalAddress:   decodeIPv4Address(tableRow.DwLocalAddr),
			LocalPort:      decodePort(tableRow.DwLocalPort),
			RemoteAddress:  decodeIPv4Address(tableRow.DwRemoteAddr),
			RemotePort:     decodePort(tableRow.DwRemotePort),
			OwnerProcessID: int(tableRow.DwOwningPid),
			State:          interop.MibTcpState(tableRow.DwState).String(),
			Kind:           "tcp4",
		}
		result = append(result, item)
		tableRowIndex += tableRowStep
	}
	return result, nil
}

func getPortBindingListForTCP6() ([]PortBindingInfo, error) {
	var table *interop.MibTcp6TableOwnerPid
	tableBufferSize := uint32(0)

	err := interop.GetExtendedTcpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET6,
		interop.TcpTableClassOwnerPidAll,
		0,
	)
	if err != nil && err != windows.ERROR_INSUFFICIENT_BUFFER {
		return nil, err
	}

	tableBuffer := make([]byte, tableBufferSize)
	table = (*interop.MibTcp6TableOwnerPid)(unsafe.Pointer(&tableBuffer[0]))
	err = interop.GetExtendedTcpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET6,
		interop.TcpTableClassOwnerPidAll,
		0,
	)
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	if table.DwNumEntries == 0 {
		return result, nil
	}

	tableRowIndex := int(unsafe.Sizeof(table.DwNumEntries))
	tableRowStep := int(unsafe.Sizeof(table.Table))
	for i := 0; i < int(table.DwNumEntries); i++ {
		tableRow := (*interop.MibTcp6RowOwnerPid)(unsafe.Pointer(&tableBuffer[tableRowIndex]))
		item := PortBindingInfo{
			LocalAddress:   decodeIPv6Address(tableRow.UcLocalAddr),
			LocalPort:      decodePort(tableRow.DwLocalPort),
			RemoteAddress:  decodeIPv6Address(tableRow.UcRemoteAddr),
			RemotePort:     decodePort(tableRow.DwRemotePort),
			OwnerProcessID: int(tableRow.DwOwningPid),
			State:          interop.MibTcpState(tableRow.DwState).String(),
			Kind:           "tcp6",
		}
		result = append(result, item)
		tableRowIndex += tableRowStep
	}
	return result, nil
}

func getPortBindingListForUDP4() ([]PortBindingInfo, error) {
	var table *interop.MibUdpTableOwnerPid
	tableBufferSize := uint32(0)

	err := interop.GetExtendedUdpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET,
		interop.UdpTableClassOwnerPid,
		0,
	)
	if err != nil && err != windows.ERROR_INSUFFICIENT_BUFFER {
		return nil, err
	}

	tableBuffer := make([]byte, tableBufferSize)
	table = (*interop.MibUdpTableOwnerPid)(unsafe.Pointer(&tableBuffer[0]))
	err = interop.GetExtendedUdpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET,
		interop.UdpTableClassOwnerPid,
		0,
	)
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	if table.DwNumEntries == 0 {
		return result, nil
	}

	tableRowIndex := int(unsafe.Sizeof(table.DwNumEntries))
	tableRowStep := int(unsafe.Sizeof(table.Table))
	for i := 0; i < int(table.DwNumEntries); i++ {
		tableRow := (*interop.MibUdpRowOwnerPid)(unsafe.Pointer(&tableBuffer[tableRowIndex]))
		item := PortBindingInfo{
			LocalAddress:   decodeIPv4Address(tableRow.DwLocalAddr),
			LocalPort:      decodePort(tableRow.DwLocalPort),
			OwnerProcessID: int(tableRow.DwOwningPid),
			State:          interop.MibTcpStateListen.String(),
			Kind:           "udp4",
		}
		result = append(result, item)
		tableRowIndex += tableRowStep
	}
	return result, nil
}

func getPortBindingListForUDP6() ([]PortBindingInfo, error) {
	var table *interop.MibUdp6TableOwnerPid
	tableBufferSize := uint32(0)

	err := interop.GetExtendedUdpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET6,
		interop.UdpTableClassOwnerPid,
		0,
	)
	if err != nil && err != windows.ERROR_INSUFFICIENT_BUFFER {
		return nil, err
	}

	tableBuffer := make([]byte, tableBufferSize)
	table = (*interop.MibUdp6TableOwnerPid)(unsafe.Pointer(&tableBuffer[0]))
	err = interop.GetExtendedUdpTable(
		uintptr(unsafe.Pointer(table)),
		&tableBufferSize,
		true,
		syscall.AF_INET6,
		interop.UdpTableClassOwnerPid,
		0,
	)
	if err != nil {
		return nil, err
	}

	result := []PortBindingInfo{}

	if table.DwNumEntries == 0 {
		return result, nil
	}

	tableRowIndex := int(unsafe.Sizeof(table.DwNumEntries))
	tableRowStep := int(unsafe.Sizeof(table.Table))
	for i := 0; i < int(table.DwNumEntries); i++ {
		tableRow := (*interop.MibUdp6RowOwnerPid)(unsafe.Pointer(&tableBuffer[tableRowIndex]))
		item := PortBindingInfo{
			LocalAddress:   decodeIPv6Address(tableRow.UcLocalAddr),
			LocalPort:      decodePort(tableRow.DwLocalPort),
			OwnerProcessID: int(tableRow.DwOwningPid),
			State:          interop.MibTcpStateListen.String(),
			Kind:           "udp6",
		}
		result = append(result, item)
		tableRowIndex += tableRowStep
	}
	return result, nil
}
