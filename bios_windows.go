package tools

import (
	"os/exec"
	"strings"
)

func BiosUUID() (string, error) {

	var uuid string

	cmd := exec.Command("wmic", "csproduct", "get", "uuid")
	bs, err := cmd.Output()

	if err != nil {
		return "", err
	}

	uuid = string(bs)
	uuid = strings.Split(uuid, "\n")[1]

	uuid = strings.TrimSpace(uuid)
	return uuid, nil
}
