package tools

import (
	"os/exec"
	"strings"
)

func BiosUUID() (string, error) {

	var uuid string

	cmd := exec.Command("bash", "-c", `system_profiler SPHardwareDataType | grep "Hardware UUID:" | awk '{print $3}'`)
	bs, err := cmd.Output()

	if err != nil {
		return "", err
	}

	uuid = string(bs)

	uuid = strings.TrimSpace(uuid)
	return uuid, nil
}
