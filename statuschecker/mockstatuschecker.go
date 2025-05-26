package statuschecker

import "fmt"

type MockStatusChecker struct {
	SimulateTimeout bool
	SimulateFail    bool
	SyncSeconds     int
}

func (m *MockStatusChecker) WaitForSync(folderID string, timeoutSeconds int) error {

	for i := 0; i < timeoutSeconds; i++ {

		if i == m.SyncSeconds && !m.SimulateFail {
			return nil
		} else if (i == m.SyncSeconds && !m.SimulateTimeout) || m.SimulateFail {
			return fmt.Errorf("mock sync error")
		}
	}

	return fmt.Errorf("mock sync timeout error")
}
