package statuschecker

import "fmt"

type MockStatusChecker struct {
	SimulateTimeout bool
	SimulateSuccess bool
	SyncSeconds     int
}

func (m *MockStatusChecker) WaitForSync(folderID string, timeoutSeconds int) error {

	for i := 0; i < timeoutSeconds; i++ {

		if i == m.SyncSeconds && m.SimulateSuccess {
			return nil
		} else if i == m.SyncSeconds && !m.SimulateTimeout && !m.SimulateSuccess {
			return fmt.Errorf("Mock sync error.")
		}
	}

	return fmt.Errorf("Mock sync timeout error.")
}
