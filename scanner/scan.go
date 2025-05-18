package scanner

import (
	"fmt"
	"os/exec"
)

func ScanFile(path string) error {
	cmd := exec.Command("clamscan", "--infected", "--recursive", "--scan-archive=yes", path)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return fmt.Errorf("virus scan failed: %v", err)
	}
	return nil
}
