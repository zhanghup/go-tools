// +build windows

package twindows

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/svc/mgr"
)

func exePath() (string, error) {
	return filepath.Abs(os.Args[0])
}

func install(name, desc string, args ...string) error {

	exepath, err := exePath()

	if err != nil {
		return err
	}

	m, err := mgr.Connect()

	if err != nil {
		return err
	}

	defer m.Disconnect()

	s, err := m.OpenService(name)

	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", name)
	}

	s, err = m.CreateService(name, exepath, mgr.Config{
		DisplayName: desc,
		StartType:   mgr.StartAutomatic}, args...)

	defer s.Close()
	return err
}

func uninstall(name string) error {

	m, err := mgr.Connect()

	if err != nil {
		return err
	}

	defer m.Disconnect()
	s, err := m.OpenService(name)

	if err != nil {
		return fmt.Errorf("service %s is not installed", name)
	}

	defer s.Close()

	err = s.Delete()
	return err
}
