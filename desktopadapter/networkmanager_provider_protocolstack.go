package desktopadapter

import (
	"fmt"
	"net"

	"github.com/AtelierNeon/an-core-api-golang/networkmanager"
)

type NetworkManagerProviderProtocolStack struct {
	networkmanager.ProtocolStack
	localAddress string
}

func (nmpps *NetworkManagerProviderProtocolStack) GetAvailablePort() (uint16, error) {
	tcpAddr, err := net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf("%s:0", nmpps.localAddress),
	)
	if err != nil {
		return 0, err
	}

	tcpListener, err := net.ListenTCP(
		"tcp",
		tcpAddr,
	)
	if err != nil {
		return 0, err
	}

	defer tcpListener.Close()
	realTcpAddr := tcpListener.Addr().(*net.TCPAddr)
	if realTcpAddr == nil {
		return 0, fmt.Errorf("failed to get real TCP address")
	}

	return uint16(realTcpAddr.Port), nil
}

func (nmpps *NetworkManagerProviderProtocolStack) IsPortAvailable(port uint16) (bool, error) {
	tcpAddr, err := net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf("%s:%d", nmpps.localAddress, port),
	)
	if err != nil {
		return false, err
	}

	tcpListener, err := net.ListenTCP(
		"tcp",
		tcpAddr,
	)
	if err != nil {
		return false, err
	}

	defer tcpListener.Close()
	return true, nil
}
