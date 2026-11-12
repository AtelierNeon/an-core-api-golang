package interop

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unsafe"
)

type ProcessID int32

func GetRunningProcessIDList() ([]int, error) {
	r1, _, errno := ProcListpids(
		processInfoTypeAllPIDs,
		0,
		nil,
		0,
	)
	kernelReturnType := KernelReturnType(errno)
	if kernelReturnType != KernelReturnTypeSuccess {
		return nil, errors.New(kernelReturnType.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to get running process ID list")
	}

	bufferSize := uint32(r1)
	buffer := make([]byte, bufferSize)
	r1, _, errno = ProcListpids(
		processInfoTypeAllPIDs,
		0,
		unsafe.Pointer(&buffer[0]),
		bufferSize,
	)
	kernelReturnType = KernelReturnType(errno)
	if kernelReturnType != KernelReturnTypeSuccess {
		return nil, errors.New(kernelReturnType.String())
	}

	if r1 == 0 {
		return nil, errors.New("failed to get running process ID list")
	}

	var pprocessID int32
	processIDListSize := int(bufferSize) / binary.Size(pprocessID)
	processIDList := make([]int, 0, processIDListSize)
	binaryBuffer := bytes.NewBuffer(buffer)

	for i := 0; i < processIDListSize; i++ {
		if readErr := binary.Read(binaryBuffer, binary.NativeEndian, &pprocessID); readErr != nil {
			return nil, readErr
		}
		if pprocessID == 0 {
			continue
		}

		processIDList = append(processIDList, int(pprocessID))
	}
	return processIDList, nil
}
