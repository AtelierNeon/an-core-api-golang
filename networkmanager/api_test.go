package networkmanager_test

import (
	"testing"

	"github.com/AtelierNeon/an-core-api-golang/networkmanager"
)

func Test_0_GetInstance(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	networkmanagerInstAlt := networkmanager.GetInstance()
	if networkmanagerInstAlt == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	if networkManagerInst != networkmanagerInstAlt {
		t.Errorf("Error instance should be the same")
	}
}

func Test_1_GetProtocolStackForEveryAddress(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	_, err := networkManagerInst.GetProtocolStackForEveryLocalAddress()
	if err != nil {
		t.Errorf("Error getter protocol stack error should be nil")
	}
}

func Test_2_GetAvailablePortForAllAddresses(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	protocolStack, err := networkManagerInst.GetProtocolStackForEveryLocalAddress()
	if err != nil {
		t.Errorf("Error getter protocol stack error should be nil")
	}
	_, err = protocolStack.GetAvailablePort()
	if err == nil {
		t.Errorf("Error getter available port error should NOT be nil")
	}
}

func Test_3_GetAvailablePortForLocalhost(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	protocolStack, err := networkManagerInst.GetProtocolStackForLocalhost()
	if err != nil {
		t.Errorf("Error getter protocol stack error should be nil")
	}
	_, err = protocolStack.GetAvailablePort()
	if err == nil {
		t.Errorf("Error getter available port error should NOT be nil")
	}
}

func Test_4_IsPortAvailableForEveryAddresses(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	protocolStack, err := networkManagerInst.GetProtocolStackForEveryLocalAddress()
	if err != nil {
		t.Errorf("Error getter protocol stack error should be nil")
	}
	_, err = protocolStack.IsPortAvailable(0)
	if err == nil {
		t.Errorf("Error getter port availability error should NOT be nil")
	}
}

func Test_5_IsPortAvailableForLocalhost(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	protocolStack, err := networkManagerInst.GetProtocolStackForLocalhost()
	if err != nil {
		t.Errorf("Error getter protocol stack error should be nil")
	}
	_, err = protocolStack.IsPortAvailable(0)
	if err == nil {
		t.Errorf("Error getter port availability error should NOT be nil")
	}
}

func Test_6_CheckPortBindingProcessIDListForEveryAddress(t *testing.T) {
	networkManagerInst := networkmanager.GetInstance()
	if networkManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	protocolStack, err := networkManagerInst.GetProtocolStackForEveryLocalAddress()
	if err != nil {
		t.Errorf("Error getter protocol stack error should be nil")
	}
	_, err = protocolStack.CheckPortBindingProcessIDList(0)
	if err == nil {
		t.Errorf("Error checking port binding error should NOT be nil")
	}
}
