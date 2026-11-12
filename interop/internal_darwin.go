package interop

import (
	"syscall"
	"unsafe"
)

const (
	procPidPathInfoSize = 0x400
)

// We should also define the relative entry in assembly (*.s)
//
//go:cgo_import_dynamic libc_proc_listpids proc_listpids "/usr/lib/libSystem.B.dylib"
var libc_proc_listpids_trampoline_addr uintptr

// We should also define the relative entry in assembly (*.s)
//
//go:cgo_import_dynamic libc_proc_pidpath proc_pidpath "/usr/lib/libSystem.B.dylib"
var libc_proc_pidpath_trampoline_addr uintptr

//go:nosplit
func proc_listpids(
	processInfoType ProcessInfoType,
	processInfo uint32,
	buffer unsafe.Pointer,
	bufferSize uint32) (r1, r2 uintptr, err syscall.Errno) {
	return syscall_syscall6(
		libc_proc_listpids_trampoline_addr,
		uintptr(processInfoType),
		uintptr(processInfo),
		uintptr(buffer),
		uintptr(bufferSize),
		0,
		0,
	)
}

//go:nosplit
func proc_pidpath(
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

// From sdk/go/src/runtime/string.go
//
//go:linkname runtime_gostring runtime.gostring
func runtime_gostring(p *byte) string

// From golang.org/x/sys/unix/syscall_darwin_libSystem.go
//
//go:noescape
//go:linkname syscall_syscall syscall.syscall
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

// From golang.org/x/sys/unix/syscall_darwin_libSystem.go
//
//go:noescape
//go:linkname syscall_syscall6 syscall.syscall6
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
