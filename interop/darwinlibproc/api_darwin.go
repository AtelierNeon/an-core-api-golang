package darwinlibproc

import (
	"bytes"
	"encoding/binary"
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

func GetRunningProcessIDList() ([]interop.ProcessID, error) {
	r1, _, errno := interop.ProcListpids(
		interop.ProcessInfoTypeAllPIDs,
		0,
		nil,
		0,
	)
	kernelReturnType := interop.KernelReturnType(errno)
	if kernelReturnType != interop.KernelReturnTypeSuccess {
		return nil, errors.New(kernelReturnType.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to get running process ID list")
	}

	bufferSize := uint32(r1)
	buffer := make([]byte, bufferSize)
	r1, _, errno = interop.ProcListpids(
		interop.ProcessInfoTypeAllPIDs,
		0,
		unsafe.Pointer(&buffer[0]),
		bufferSize,
	)
	kernelReturnType = interop.KernelReturnType(errno)
	if kernelReturnType != interop.KernelReturnTypeSuccess {
		return nil, errors.New(kernelReturnType.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to get running process ID list")
	}

	var pprocessID int32
	processIDListSize := int(bufferSize) / binary.Size(pprocessID)
	processIDList := make([]interop.ProcessID, 0, processIDListSize)
	binaryBuffer := bytes.NewBuffer(buffer)

	for i := 0; i < processIDListSize; i++ {
		if readErr := binary.Read(binaryBuffer, binary.NativeEndian, &pprocessID); readErr != nil {
			return nil, readErr
		}
		if pprocessID == 0 {
			continue
		}

		processIDList = append(processIDList, interop.ProcessID(pprocessID))
	}
	return processIDList, nil
}
