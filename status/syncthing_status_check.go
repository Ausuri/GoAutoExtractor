package statuscheck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func WaitForSync(folderID string) error {

	baseURL := os.Getenv("ST_BASE_URL")
	url := fmt.Sprintf("%s/rest/db/status?folder=%s", baseURL, folderID)

	for {
		apiKey := os.Getenv("ST_API_KEY")
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
