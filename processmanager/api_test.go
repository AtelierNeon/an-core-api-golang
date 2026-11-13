package processmanager_test

import (
	"testing"

	"github.com/AtelierNeon/an-core-api-golang/processmanager"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	// Something to setup
	//logger.GetInstance().SetSeverity(logger.SeverityTypeInfo)

	return func(tb testing.TB) {
		// Something to teardown
	}
}

func Test_0_GetInstance(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	processManagerInstAlt := processmanager.GetInstance()
	if processManagerInstAlt == nil {
		t.Errorf("Error instance should not be nil")
	}
	if processManagerInst != processManagerInstAlt {
		t.Errorf("Error instance should be the same")
	}
}

func Test_1_GetProcessIdListByName(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	_, err := processManagerInst.GetProcessIDListByName("AAA.exe")
	if err == nil {
		t.Errorf("Error getter error should not be nil")
	}
}

func Test_2_GetProcessNameById(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	_, err := processManagerInst.GetProcessNameByID(0)
	if err == nil {
		t.Errorf("Error getter error should not be nil")
	}
}

func Test_3_GetProcessPathById(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	_, err := processManagerInst.GetProcessPathByID(0)
	if err == nil {
		t.Errorf("Error getter error should not be nil")
	}
}
