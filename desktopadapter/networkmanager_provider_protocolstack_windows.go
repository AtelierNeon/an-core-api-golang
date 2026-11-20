package desktopadapter

import (
	"net"

	"github.com/AtelierNeon/an-core-api-golang/interop"
	"github.com/AtelierNeon/an-core-api-golang/interop/windowsmib"
)

func (nmpps *NetworkManagerProviderProtocolStack) CheckPortBindingProcessIDList(port uint16) ([]int, error) {
	portBindingList, err := windowsmib.GetPortBindingList()
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
		if ipPortBindingLocalAddress == nil {
			continue
		}

		if portBinding.Kind == interop.PortBindingKindTCP4 || portBinding.Kind == interop.PortBindingKindUDP4 {
			if localAddress == "localhost" {
				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Localhost) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}

				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Any) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}
			}

			ipLocalAddress := net.ParseIP(localAddress)
			if ipLocalAddress == nil {
				continue
			}
			if ipLocalAddress.To4() == nil { // Skip non-IPv4 addresses for IPv4 bindings
				continue
			}

			if ipLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Localhost) {
				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Localhost) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}

				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Any) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}
			}

			if ipPortBindingLocalAddress.Equal(ipLocalAddress) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV4Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}

		if portBinding.Kind == interop.PortBindingKindTCP6 || portBinding.Kind == interop.PortBindingKindUDP6 {
			if localAddress == "localhost" {
				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Localhost) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}

				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Any) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}
			}

			ipLocalAddress := net.ParseIP(localAddress)
			if ipLocalAddress == nil {
				continue
			}
			if ipLocalAddress.To4() != nil { // Skip IPv4 addresses for IPv6 bindings
				continue
			}

			if ipLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Localhost) {
				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Localhost) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}

				if ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Any) {
					result = append(result, int(portBinding.OwnerProcessID))
					continue
				}
			}

			if ipPortBindingLocalAddress.Equal(ipLocalAddress) || ipPortBindingLocalAddress.Equal(networkManagerProviderProtocolStack_ipV6Any) {
				result = append(result, int(portBinding.OwnerProcessID))
				continue
			}
		}
	}
	return result, nil
}
