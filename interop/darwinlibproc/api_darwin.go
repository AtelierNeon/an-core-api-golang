package darwinlibproc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
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

	if r1 == 0 {
		return "", errors.New("invalid process path size")
	}

	return interop.BufferToString(buffer), nil
}

func ListRunningProcessIDs() ([]interop.ProcessID, error) {
	r1, _, errno := interop.ProcListpids(
		interop.ProcessInfoTypeAllPIDs,
		0,
		nil,
		0,
	)
	kernelReturn := interop.KernelReturnType(errno)
	if kernelReturn != interop.KernelReturnTypeSuccess {
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
	r1, _, errno = interop.ProcListpids(
		interop.ProcessInfoTypeAllPIDs,
		0,
		unsafe.Pointer(&buffer[0]),
		bufferSize,
	)
	kernelReturn = interop.KernelReturnType(errno)
	if kernelReturn != interop.KernelReturnTypeSuccess {
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
