package interop

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	ANY_SIZE = 1
)

const (
	NTFS_MAX_PATH = 32768
)

var (
	modiphlpapi             = windows.NewLazySystemDLL("iphlpapi.dll")
	procGetExtendedTcpTable = modiphlpapi.NewProc("GetExtendedTcpTable")
	procGetExtendedUdpTable = modiphlpapi.NewProc("GetExtendedUdpTable")
)

type ProcessID uint32

type MibTcp6RowOwnerPid struct {
	UcLocalAddr     [16]byte
	DwLocalScopeId  uint32
	DwLocalPort     uint32
	UcRemoteAddr    [16]byte
	DwRemoteScopeId uint32
	DwRemotePort    uint32
	DwState         uint32
	DwOwningPid     uint32
}

type MibTcpRowOwnerPid struct {
	DwState      uint32
	DwLocalAddr  uint32
	DwLocalPort  uint32
	DwRemoteAddr uint32
	DwRemotePort uint32
	DwOwningPid  uint32
}

type MibTcp6TableOwnerPid struct {
	DwNumEntries uint32
	Table        [ANY_SIZE]MibTcp6RowOwnerPid
}

type MibTcpTableOwnerPid struct {
	DwNumEntries uint32
	Table        [ANY_SIZE]MibTcpRowOwnerPid
}

type MibUdp6RowOwnerPid struct {
	UcLocalAddr    [16]byte
	DwLocalScopeId uint32
	DwLocalPort    uint32
	DwOwningPid    uint32
}

type MibUdpRowOwnerPid struct {
	DwLocalAddr uint32
	DwLocalPort uint32
	DwOwningPid uint32
}

type MibUdp6TableOwnerPid struct {
	DwNumEntries uint32
	Table        [ANY_SIZE]MibUdp6RowOwnerPid
}

type MibUdpTableOwnerPid struct {
	DwNumEntries uint32
	Table        [ANY_SIZE]MibUdpRowOwnerPid
}

func GetExtendedTcpTable(
	pTcpTable uintptr,
	pdwSize *uint32,
	bOrder bool,
	ulAf uint32,
	tableClass TcpTableClass,
	reserved uint32) error {
	r1, _, _ := syscall.SyscallN(
		procGetExtendedTcpTable.Addr(),
		pTcpTable,
		uintptr(unsafe.Pointer(pdwSize)),
		ToUintptrFromBool(bOrder),
		uintptr(ulAf),
		uintptr(tableClass),
		uintptr(reserved),
	)
	if r1 != 0 {
		return syscall.Errno(r1)
	}

	return nil
}

func ToUintptrFromBool(value bool) uintptr {
	if value {
		return 1
	}

	return 0
}

func GetExtendedUdpTable(
	pUdpTable uintptr,
	pdwSize *uint32,
	bOrder bool,
	ulAf uint32,
	tableClass UdpTableClass,
	reserved uint32) error {
	r1, _, _ := syscall.SyscallN(
		procGetExtendedUdpTable.Addr(),
		pUdpTable,
		uintptr(unsafe.Pointer(pdwSize)),
		ToUintptrFromBool(bOrder),
		uintptr(ulAf),
		uintptr(tableClass),
		uintptr(reserved),
	)
	if r1 != 0 {
		return syscall.Errno(r1)
	}

	return nil
}
