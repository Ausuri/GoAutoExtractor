package statuscheck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiKey   = "YOUR_API_KEY"
	baseURL  = "http://localhost:8384"
)

func WaitForSync(folderID string) error {
	url := fmt.Sprintf("%s/rest/db/status?folder=%s", baseURL, folderID)

	for {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("X-API-Key", apiKey)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var data map[string]interface{}
		json.Unmarshal(body, &data)

		state := data["state"].(string)
		needBytes := data["needBytes"].(float64)

		if state == "idle" && needBytes == 0 {
			fmt.Println("✅ Synced.")
			return nil
		}

		fmt.Printf("⏳ Syncing... (%s, needBytes=%.0f)\n", state, needBytes)
		time.Sleep(2 * time.Second)
	}
}
