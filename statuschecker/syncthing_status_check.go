package statuschecker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type SyncthingStatusChecker struct{}

const TIMER_SECONDS_INTERVAL = 1

func (s SyncthingStatusChecker) WaitForSync(folderID string, timeoutSeconds int) error {

	baseURL := os.Getenv("ST_BASE_URL")
	url := fmt.Sprintf("%s/rest/db/status?folder=%s", baseURL, folderID)
	secondsCounter := 0
	apiKey := os.Getenv("ST_API_KEY")

	for secondsCounter < timeoutSeconds {
		isSynced, err := pollStatus(url, apiKey)
		secondsCounter += TIMER_SECONDS_INTERVAL
		time.Sleep(TIMER_SECONDS_INTERVAL * time.Second)

		if isSynced && err == nil {
			return nil
		}

		if err != nil {
			return err
		}
	}

	return fmt.Errorf("timeout waiting for sync to finish.")
}

func pollStatus(url string, apiKey string) (isSynced bool, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	//Parse the response body
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	state := data["state"].(string)
	needBytes := data["needBytes"].(float64)

	//If the state is idle and needBytes is 0, the folder is synced.
	if state == "idle" && needBytes == 0 {
		fmt.Println("Synced.")
		return true, nil
	}

	updateProgress(state, needBytes)
	return false, nil
}

func updateProgress(state string, needBytes float64) {
	fmt.Printf("Syncing... (%s, needBytes=%.0f)\n", state, needBytes)
}
