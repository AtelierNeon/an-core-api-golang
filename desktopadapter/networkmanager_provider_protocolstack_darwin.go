package desktopadapter

import (
	"net"

	"github.com/AtelierNeon/an-core-api-golang/interop/darwinlibproc"
)

func (nmpps *NetworkManagerProviderProtocolStack) CheckPortBindingProcessIDList(port uint16) ([]int, error) {
	portBindingList, err := darwinlibproc.GetPortBindingList()
	if err != nil {
		return nil, err
	}

	var result []int
	for _, portBinding := range portBindingList {
		if portBinding.LocalPort != port {
			continue
		}

		localAddress := nmpps.localAddress

		if localAddress == "" {
			result = append(result, int(portBinding.OwnerProcessID))
			continue
		}

		ipPortBindingLocalAddress := net.ParseIP(portBinding.LocalAddress)

		if localAddress == "localhost" {
			if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Localhost) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Localhost) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}

			if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Any) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}

		ipLocalAddress := net.ParseIP(localAddress)

		if ipLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Localhost) {
			if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Localhost) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}

		if ipLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Localhost) {
			if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Localhost) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}

		if ipLocalAddress.To4() != nil {
			if ipPortBindingLocalAddress.Equal(ipLocalAddress) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		} else {
			if ipPortBindingLocalAddress.Equal(ipLocalAddress) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}
	}
	return result, nil
}
