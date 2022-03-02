package tools

import (
	"os"
	"os/exec"
	"strings"
)

func BiosUUID() (string, error) {

	var uuid string

	var err error

	err = (func() error {

		files := []string{"/sys/class/dmi/id/product_uuid", "/var/lib/dbus/machine-id", "/etc/machine-id"}

		var lastError error

		for _, file := range files {

			cmd := exec.Command("cat", file)
			bs, err := cmd.Output()

			if err != nil {
				lastError = err
			}

			if len(bs) > 0 {

				uuid = string(bs)
				return nil
			}
		}

		return lastError
	})()

	// alternative method for windows WSL
	if err != nil {

		WMIC := "/mnt/c/Windows/System32/wbem/wmic.exe"

		info, err := os.Stat(WMIC)

		if os.IsNotExist(err) || info.IsDir() {
			return "", err
		}

		cmd := exec.Command(WMIC, "csproduct", "get", "uuid")
		bs, err := cmd.Output()

		if err != nil {
			return "", err
		}

		uuid = string(bs)
		uuid = strings.Split(uuid, "\n")[1]
	}

	uuid = strings.TrimSpace(uuid)
	return uuid, nil
}
