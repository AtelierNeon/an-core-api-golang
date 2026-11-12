package darwinlibproc

import (
	"errors"
	"unsafe"

	"github.com/AtelierNeon/an-core-api-golang/interop"
)

func GetProcessPathByProcessID(processID interop.ProcessID) (string, error) {
	buffer := interop.NewProcPidpathBuffer()
	r1, _, errno := interop.ProcPidpath(
		processID,
		unsafe.Pointer(&buffer[0]),
		uint32(len(buffer)),
	)

	kernelReturn := interop.KernelReturnType(errno)
	if kernelReturn != interop.KernelReturnTypeSuccess {
		return "", errors.New(kernelReturn.String())
	}

	if r1 <= 0 {
		return "", errors.New("invalid process path size")
	}

	return interop.BufferToString(buffer), nil
}
