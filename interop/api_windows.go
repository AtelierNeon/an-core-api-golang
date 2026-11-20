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

const (
	FALSE = WindowsBool(0)
	TRUE  = WindowsBool(1)
)

var (
	modiphlpapi             = windows.NewLazySystemDLL("iphlpapi.dll")
	procGetExtendedTcpTable = modiphlpapi.NewProc("GetExtendedTcpTable")
	procGetExtendedUdpTable = modiphlpapi.NewProc("GetExtendedUdpTable")
)

type ProcessID WindowsDWord

type WindowsBool int32

type WindowsEnum int32

type WindowsDWord uint32

type WindowsUChar uint8

type WindowsULong uint32

/* https://learn.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcp6row_owner_pid */
type MIB_TCP6ROW_OWNER_PID struct {
	UcLocalAddr     [16]WindowsUChar
	DwLocalScopeId  WindowsDWord
	DwLocalPort     WindowsDWord
	UcRemoteAddr    [16]WindowsUChar
	DwRemoteScopeId WindowsDWord
	DwRemotePort    WindowsDWord
	DwState         WindowsDWord
	DwOwningPid     ProcessID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcp6table_owner_pid */
type MIB_TCP6TABLE_OWNER_PID struct {
	DwNumEntries WindowsDWord
	Table        [ANY_SIZE]MIB_TCP6ROW_OWNER_PID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcprow_owner_pid */
type MIB_TCPROW_OWNER_PID struct {
	DwState      WindowsDWord
	DwLocalAddr  WindowsDWord
	DwLocalPort  WindowsDWord
	DwRemoteAddr WindowsDWord
	DwRemotePort WindowsDWord
	DwOwningPid  ProcessID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/tcpmib/ns-tcpmib-mib_tcptable_owner_pid */
type MIB_TCPTABLE_OWNER_PID struct {
	DwNumEntries WindowsDWord
	Table        [ANY_SIZE]MIB_TCPROW_OWNER_PID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/udpmib/ns-udpmib-mib_udp6row_owner_pid */
type MIB_UDP6ROW_OWNER_PID struct {
	UcLocalAddr    [16]WindowsUChar
	DwLocalScopeId WindowsDWord
	DwLocalPort    WindowsDWord
	DwOwningPid    ProcessID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/udpmib/ns-udpmib-mib_udp6table_owner_pid */
type MIB_UDP6TABLE_OWNER_PID struct {
	DwNumEntries WindowsDWord
	Table        [ANY_SIZE]MIB_UDP6ROW_OWNER_PID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/udpmib/ns-udpmib-mib_udprow_owner_pid */
type MIB_UDPROW_OWNER_PID struct {
	DwLocalAddr WindowsDWord
	DwLocalPort WindowsDWord
	DwOwningPid ProcessID
}

/* https://learn.microsoft.com/en-us/windows/win32/api/udpmib/ns-udpmib-mib_udptable_owner_pid */
type MIB_UDPTABLE_OWNER_PID struct {
	DwNumEntries WindowsDWord
	Table        [ANY_SIZE]MIB_UDPROW_OWNER_PID
}

func GetExtendedTcpTable(
	pTcpTable uintptr,
	pdwSize *WindowsDWord,
	bOrder WindowsBool,
	ulAf WindowsULong,
	tableClass TcpTableClass,
	reserved WindowsULong) error {
	r1, _, _ := syscall.SyscallN(
		procGetExtendedTcpTable.Addr(),
		pTcpTable,
		uintptr(unsafe.Pointer(pdwSize)),
		uintptr(bOrder),
		uintptr(ulAf),
		uintptr(tableClass),
		uintptr(reserved),
	)
	if r1 != 0 {
		return syscall.Errno(r1)
	}

	return nil
}

func GetExtendedUdpTable(
	pUdpTable uintptr,
	pdwSize *WindowsDWord,
	bOrder WindowsBool,
	ulAf WindowsULong,
	tableClass UdpTableClass,
	reserved WindowsULong) error {
	r1, _, _ := syscall.SyscallN(
		procGetExtendedUdpTable.Addr(),
		pUdpTable,
		uintptr(unsafe.Pointer(pdwSize)),
		uintptr(bOrder),
		uintptr(ulAf),
		uintptr(tableClass),
		uintptr(reserved),
	)
	if r1 != 0 {
		return syscall.Errno(r1)
	}

	return nil
}
