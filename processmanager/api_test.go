package processmanager_test

import (
	"testing"

	"github.com/AtelierNeon/an-core-api-golang/processmanager"
)

func Test_0_GetInstance(t *testing.T) {
	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	processManagerInstAlt := processmanager.GetInstance()
	if processManagerInstAlt == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	if processManagerInst != processManagerInstAlt {
		t.Errorf("Error instance should be the same")
	}
}

func Test_1_GetProcessIdListByName(t *testing.T) {
	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	_, err := processManagerInst.GetProcessIDListByName("AAA.exe")
	if err == nil {
		t.Errorf("Error getter error should NOT be nil")
	}
}

func Test_2_GetProcessNameById(t *testing.T) {
	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	_, err := processManagerInst.GetProcessNameByID(0)
	if err == nil {
		t.Errorf("Error getter error should NOT be nil")
	}
}

func Test_3_GetProcessPathById(t *testing.T) {
	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should NOT be nil")
	}
	_, err := processManagerInst.GetProcessPathByID(0)
	if err == nil {
		t.Errorf("Error getter error should NOT be nil")
	}
}
