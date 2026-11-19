package interop

type LinuxPidT int32

type ProcessID LinuxPidT

func GetInitProcessID() ProcessID {
	return 1
}
