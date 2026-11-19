package desktopadapter

import "github.com/AtelierNeon/an-core-api-golang/interop/linuxprocfs"

func (nmpps *NetworkManagerProviderProtocolStack) CheckPortBindingProcessID(port uint16) (int, error) {
	portBindingList, err := linuxprocfs.GetPortBindingList()
	if err != nil {
		return -1, err
	}

	for _, portBinding := range portBindingList {
		if portBinding.LocalPort == port {
			localAddress := nmpps.localAddress

			if localAddress == "" {
				return portBinding.OwnerProcessID, nil
			}

			if localAddress == "localhost" {
				if portBinding.LocalAddress == "127.0.0.1" || portBinding.LocalAddress == "::1" {
					return portBinding.OwnerProcessID, nil
				}
			}

			if portBinding.LocalAddress == localAddress {
				return portBinding.OwnerProcessID, nil
			}
		}
	}
	return -1, nil
}
