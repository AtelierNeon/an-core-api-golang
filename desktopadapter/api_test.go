package desktopadapter_test

import (
	"reflect"
	"testing"

	"github.com/AtelierNeon/an-core-api-golang/desktopadapter"
	"github.com/AtelierNeon/an-core-api-golang/logger"
	"github.com/AtelierNeon/an-core-api-golang/processmanager"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	// Something to setup
	//logger.GetInstance().SetSeverity(logger.SeverityTypeInfo)

	config := desktopadapter.NewConfig()
	desktopadapter.ConfigureWithConfig(config)

	return func(tb testing.TB) {
		// Something to teardown
	}
}

func Test_logger_0_GetInstance(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInstAlt := logger.GetInstance()
	if loggerInstAlt == nil {
		t.Errorf("Error instance should not be nil")
	}
	if loggerInst != loggerInstAlt {
		t.Errorf("Error instance should be the same")
	}
}

func Test_logger_0_GetInstance_WithName(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInstAlt := logger.GetInstanceWithLabel("Alt")
	if loggerInstAlt == nil {
		t.Errorf("Error instance should not be nil")
	}
	if loggerInst == loggerInstAlt {
		t.Errorf("Error instance should not be the same")
	}
}

func Test_logger_0_GetInstance_WithType(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInstAlt := logger.GetInstanceWithType(reflect.TypeOf("Alt"))
	if loggerInstAlt == nil {
		t.Errorf("Error instance should not be nil")
	}
	if loggerInst == loggerInstAlt {
		t.Errorf("Error instance should not be the same")
	}
}

func Test_logger_1_Debug(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInst.Debug("Default test debug message 1")
	loggerInst.DebugF("Default test debug message %d", 2)
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
}

func Test_logger_2_Error(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInst.Error("Default test error message 1")
	loggerInst.ErrorF("Default test error message %d", 2)
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
}

func Test_logger_3_Fatal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping in short testing")
	}

	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInst.Fatal("Default test fatal message 1")
	loggerInst.FatalF("Default test fatal message %d", 2)
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
}

func Test_logger_4_Info(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInst.Info("Default test info message 1")
	loggerInst.InfoF("Default test info message %d", 2)
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
}

func Test_logger_5_Trace(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInst.Trace("Default test trace message 1")
	loggerInst.TraceF("Default test trace message %d", 2)
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
}

func Test_logger_6_Warn(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	loggerInst := logger.GetInstance()
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	loggerInst.Warn("Default test warn message 1")
	loggerInst.WarnF("Default test warn message %d", 2)
	if loggerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
}

func Test_processmanager_0_GetInstance(t *testing.T) {
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

func Test_processmanager_1_GetProcessIdListByName(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	if processManagerInst.IsCurrentProcessRunningInDockerContainer() {
		return
	}

	var processNameRunning = getRunningProcessNameToTest()
	processIdListRunning, err := processManagerInst.GetProcessIDListByName(processNameRunning)
	if err != nil {
		t.Errorf("Error getter error should be nil")
	}
	if len(processIdListRunning) <= 0 {
		t.Errorf("Error running process id list size should be greater than zero")
	}
	var processNameNotRunning = "NotRunningProcess.exe"
	processIdListNotRunning, err := processManagerInst.GetProcessIDListByName(processNameNotRunning)
	if err != nil {
		t.Errorf("Error getter error should be nil")
	}
	if len(processIdListNotRunning) > 0 {
		t.Errorf("Error not-running process id list size should be zero")
	}
}

func Test_processmanager_2_GetProcessNameById(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	if processManagerInst.IsCurrentProcessRunningInDockerContainer() {
		return
	}

	var processNameRunning = getRunningProcessNameToTest()
	processIdListRunning, err := processManagerInst.GetProcessIDListByName(processNameRunning)
	if err != nil {
		t.Errorf("Error getter error should be nil")
	}
	if len(processIdListRunning) <= 0 {
		t.Errorf("Error running process id list size should be greater than zero")
	}
	processIdRunning := processIdListRunning[0]
	processName, err := processManagerInst.GetProcessNameByID(processIdRunning)
	if err != nil {
		t.Errorf("Error getter error should be nil")
	}
	if processName == "" {
		t.Errorf("Error running process name should not be empty")
	}
	if processName != processNameRunning {
		t.Errorf("Error running process name should not the same")
	}
}

func Test_processmanager_3_GetProcessPathById(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	processManagerInst := processmanager.GetInstance()
	if processManagerInst == nil {
		t.Errorf("Error instance should not be nil")
	}
	if processManagerInst.IsCurrentProcessRunningInDockerContainer() {
		return
	}

	var processNameRunning = getRunningProcessNameToTest()
	processIdListRunning, err := processManagerInst.GetProcessIDListByName(processNameRunning)
	if err != nil {
		t.Errorf("Error getter error should be nil")
	}
	if len(processIdListRunning) <= 0 {
		t.Errorf("Error running process id list size should be greater than zero")
	}
	processIdRunning := processIdListRunning[0]
	processName, err := processManagerInst.GetProcessPathByID(processIdRunning)
	if err != nil {
		t.Errorf("Error getter error should be nil")
	}
	if processName == "" {
		t.Errorf("Error running process path should not be empty")
	}
}
