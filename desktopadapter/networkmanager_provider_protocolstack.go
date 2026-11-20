package desktopadapter

import (
	"fmt"
	"net"
	"strconv"

	"github.com/AtelierNeon/an-core-api-golang/networkmanager"
)

var (
	networkManagerProviderProtocolStack_ipV4Any       = net.ParseIP("0.0.0.0")
	networkManagerProviderProtocolStack_ipV4Localhost = net.ParseIP("127.0.0.1")
	networkManagerProviderProtocolStack_ipV6Any       = net.ParseIP("::")
	networkManagerProviderProtocolStack_ipV6Localhost = net.ParseIP("::1")
)

type NetworkManagerProviderProtocolStack struct {
	networkmanager.ProtocolStack
	localAddress string
}

func (nmpps *NetworkManagerProviderProtocolStack) GetAvailablePort() (uint16, error) {
	tcpAddr, err := net.ResolveTCPAddr(
		"tcp",
		net.JoinHostPort(nmpps.localAddress, "0"),
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
		net.JoinHostPort(nmpps.localAddress, strconv.Itoa(int(port))),
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
