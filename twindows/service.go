package twindows

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/svc"
)

type anyservice struct {
	main func(...string)
	args []string
}

func chdir(exe string) error {

	abs, err := filepath.Abs(filepath.Dir(exe))

	if err != nil {
		return err
	}

	return os.Chdir(abs)
}

func (m *anyservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue

	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	err := chdir(m.args[0])

	if err != nil {
		panic(err)
	}

	var cmd *exec.Cmd

	if len(m.args) > 1 {
		cmd = exec.Command(m.args[0], m.args[1:]...)
	} else {
		cmd = exec.Command(m.args[0])
	}

	go func() {

		stderr, err := cmd.StderrPipe()

		if err != nil {
			ioutil.WriteFile("./error.log", []byte(err.Error()), 0644)
			return
		}

		fp, err := os.OpenFile("./error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

		if err != nil {
			ioutil.WriteFile("./error.log", []byte(err.Error()), 0644)
			return
		}

		_, err = io.Copy(fp, stderr)

		if err != nil {
			ioutil.WriteFile("./error.log", []byte(err.Error()), 0644)
			return
		}
	}()

	err = cmd.Start()

	if err != nil {
		panic(err)
	}

	die := make(chan bool)
	go func() {
		cmd.Wait()
		die <- true
	}()

	for {
		select {
		case <-die:
			changes <- svc.Status{State: svc.Stopped}
			os.Exit(0)
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
			case svc.Stop, svc.Shutdown:
				changes <- svc.Status{State: svc.StopPending}
				cmd.Process.Kill()
				changes <- svc.Status{State: svc.Stopped}
				time.Sleep(2 * time.Second)
				os.Exit(0)
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
			}
		}
	}
}
