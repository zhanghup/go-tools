// +build windows

package twindows

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func start(name string) error {

	m, err := mgr.Connect()

	if err != nil {
		return err
	}

	defer m.Disconnect()
	s, err := m.OpenService(name)

	if err != nil {
		panic(err)
	}

	defer s.Close()
	err = s.Start()

	if err != nil {
		panic(err)
	}

	return nil
}

func control(name string, c svc.Cmd, to svc.State) error {

	m, err := mgr.Connect()

	if err != nil {
		return err
	}

	defer m.Disconnect()
	s, err := m.OpenService(name)

	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}

	defer s.Close()

	status, err := s.Control(c)

	if err != nil {
		return fmt.Errorf("could not send control=%d: %v", c, err)
	}

	timeout := time.Now().Add(10 * time.Second)

	for status.State != to {

		if timeout.Before(time.Now()) {
			return fmt.Errorf("timeout waiting for service to go to state=%d", to)
		}

		time.Sleep(300 * time.Millisecond)

		status, err = s.Query()

		if err != nil {
			return fmt.Errorf("could not retrieve service status: %v", err)
		}
	}

	return nil
}
