package interop

import (
	"math"
	"syscall"
	"unsafe"
)

type DarwinGidT uint32

type DarwinInt int32

type DarwinOffT int64

type DarwinShort int16

type DarwinUChar uint8

type DarwinUidT uint32

type DarwinUShort uint16

type ProcessID DarwinInt

const (
	MaxDarwinInt = math.MaxInt32
)

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type In4In6_addr struct /* struct in4in6_addr */ {
	Pad   [3]uint32
	Addr4 [4]byte
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type In_sockinfo struct /* struct in_sockinfo */ {
	Fport  int32
	Lport  int32
	Gencnt uint64
	Flags  uint32
	Flow   uint32
	Vflag  uint8
	Ip_ttl uint8
	Rfu_1  uint32
	Faddr  [16]byte
	Laddr  [16]byte
	V4     In_sockinfo_V4
	V6     In_sockinfo_V6
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Proc_fdinfo struct /* struct proc_fdinfo */ {
	ProcFd     int32
	ProcFdType ProcessFdType
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Proc_fileinfo struct /* struct proc_fileinfo */ {
	OpenFlags  uint32
	Status     uint32
	Offset     DarwinOffT
	Type       int32
	GuardFlags uint32
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Socket_info struct /* struct socket_info */ {
	Stat     Vinfo_stat
	So       uint64
	Pcb      uint64
	Type     DarwinInt
	Protocol DarwinInt
	Family   AddressFamily
	Options  DarwinShort
	Linger   DarwinShort
	State    DarwinShort
	Qlen     DarwinShort
	Incqlen  DarwinShort
	Qlimit   DarwinShort
	Timeo    DarwinShort
	Error    DarwinUShort
	Oobmark  uint32
	Rcv      Sockbuff_info
	Snd      Sockbuff_info
	Kind     SocketKind
	Rfu_1    uint32
	Proto    [528]byte
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Sockbuff_info struct /* struct sockbuf_info */ {
	Cc    uint32
	Hiwat uint32
	Mbcnt uint32
	Mbmax uint32
	Lowat uint32
	Flags DarwinShort
	Timeo DarwinShort
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Socket_fdinfo struct /* struct socket_fdinfo */ {
	File   Proc_fileinfo
	Socket Socket_info
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Tcp_sockinfo struct /* struct tcp_sockinfo */ {
	Ini   In_sockinfo
	State TcpState
	Timer [4]DarwinInt
	Mss   DarwinInt
	Flags uint32
	Rfu_1 uint32
	Tp    uint64
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type In_sockinfo_V4 struct /* struct in_sockinfo / struct v4 */ {
	Tos DarwinUChar
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type In_sockinfo_V6 struct /* struct in_sockinfo / struct v6 */ {
	Hlim    uint8
	Cksum   DarwinInt
	Ifindex DarwinUShort
	Hops    DarwinShort
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/bsd/sys/proc_info.h */
type Vinfo_stat struct /* struct vinfo_stat */ {
	Dev           uint32
	Mode          uint16
	Nlink         uint16
	Ino           uint64
	Uid           DarwinUidT
	Gid           DarwinGidT
	Atime         int64
	AtimeNsec     int64
	Mtime         int64
	MtimeNsec     int64
	Ctime         int64
	CtimeNsec     int64
	Birthtime     int64
	BirthtimeNsec int64
	Size          DarwinOffT
	Blocks        int64
	Blksize       int32
	Flags         uint32
	Gen           uint32
	Rdev          uint32
	Qspare        [2]int64
}

func BufferToString(buffer []byte) string {
	return runtime_gostring(&buffer[0])
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/libsyscall/wrappers/libproc/libproc.h */
//go:nosplit
func Proc_listpids(
	processIDListType ProcessIDListType,
	processTypeInfo uint32,
	buffer unsafe.Pointer,
	bufferSize DarwinInt) (r1, r2 uintptr, err syscall.Errno) {
	return syscall_syscall6(
		libc_proc_listpids_trampoline_addr,
		uintptr(processIDListType),
		uintptr(processTypeInfo),
		uintptr(buffer),
		uintptr(bufferSize),
		0,
		0,
	)
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/libsyscall/wrappers/libproc/libproc.h */
//go:nosplit
func Proc_pidfdinfo(
	processID ProcessID,
	processFd int32,
	processFDFlavor ProcessFDFlavor,
	buffer unsafe.Pointer,
	buffersize uint32) (r1, r2 uintptr, err syscall.Errno) {
	return syscall_syscall6(
		libc_proc_pidfdinfo_trampoline_addr,
		uintptr(processID),
		uintptr(processFd),
		uintptr(processFDFlavor),
		uintptr(buffer),
		uintptr(buffersize),
		0,
	)
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/libsyscall/wrappers/libproc/libproc.h */
//go:nosplit
func Proc_pidinfo(
	processID ProcessID,
	processIDFlavor ProcessIDFlavor,
	arg uint64,
	buffer unsafe.Pointer,
	buffersize uint32) (r1, r2 uintptr, err syscall.Errno) {
	return syscall_syscall6(
		libc_proc_pidinfo_trampoline_addr,
		uintptr(processID),
		uintptr(processIDFlavor),
		uintptr(arg),
		uintptr(buffer),
		uintptr(buffersize),
		0,
	)
}

/* https://github.com/apple-oss-distributions/xnu/blob/main/libsyscall/wrappers/libproc/libproc.h */
//go:nosplit
func Proc_pidpath(
	processID ProcessID,
	buffer unsafe.Pointer,
	bufferSize uint32) (r1, r2 uintptr, err syscall.Errno) {
	return syscall_syscall(
		libc_proc_pidpath_trampoline_addr,
		uintptr(processID),
		uintptr(buffer),
		uintptr(bufferSize),
	)
}
