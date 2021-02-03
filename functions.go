package GoIPerf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

// Download runs an IPerf3 download test
func Download(server, port string) (result *Result, err error) {
	cmd := exec.Command(Location, "-J", "-R", "-c", server, "-p", port)
	out, err := cmd.CombinedOutput()

	err = json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	if result.Error != "" {
		return result, errors.New(result.Error)
	}
	return
}

// Upload runs an IPerf3 upload test
func Upload(server, port string) (result *Result, err error) {
	cmd := exec.Command(Location, "-J", "-c", server, "-p", port)
	out, err := cmd.CombinedOutput()

	err = json.Unmarshal(out, &result)
	if err != nil {
		return
	}

	if result.Error != "" {
		return result, errors.New(result.Error)
	}
	return
}
