package desktopadapter

import "github.com/AtelierNeon/an-core-api-golang/interop/windowsmib"

func (nmpps *NetworkManagerProviderProtocolStack) CheckPortBindingProcessIDList(port uint16) ([]int, error) {
	portBindingList, err := windowsmib.GetPortBindingList()
	if err != nil {
		return nil, err
	}

	var result []int
	for _, portBinding := range portBindingList {
		if portBinding.LocalPort == port {
			localAddress := nmpps.localAddress

			if localAddress == "" {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}

			if localAddress == "localhost" {
				if portBinding.LocalAddress == "127.0.0.1" || portBinding.LocalAddress == "::1" {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}

				if portBinding.LocalAddress == "0.0.0.0" || portBinding.LocalAddress == "::" {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}
			}

			if portBinding.LocalAddress == localAddress {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}
	}
	return result, nil
}
