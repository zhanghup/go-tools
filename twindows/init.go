// +build windows

package twindows

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/sys/windows/svc"
)

func run(name string, entry svc.Handler) error {
	return svc.Run(name, entry)
}

func Service(cmd, svcName string, arg2s ...string) {
	isIntSess, err := svc.IsAnInteractiveSession()

	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
		os.Exit(0)
	}

	if !isIntSess {

		args := []string{os.Args[0]}
		if len(arg2s) > 0 {
			args = append(args, arg2s...)
		}

		err := run(svcName, &anyservice{args: args})

		if err != nil {
			log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
			os.Exit(0)
		}

		select {
		case <-time.After(24 * 365 * 10 * time.Hour):
			os.Exit(0)
		}
	}

	switch cmd {
	case "install":

		args := []string{"service", "deamon", svcName}

		if len(arg2s) > 0 {
			args = append(args, arg2s...)
		}

		err = install(svcName, svcName, args...)

		if err == nil {
			fmt.Println("安装服务成功...")
			os.Exit(0)
		}

	case "uninstall":

		err = uninstall(svcName)

		if err == nil {
			fmt.Println("删除服务成功...")
			os.Exit(0)
		}

	case "start":

		err = start(svcName)

		if err == nil {
			fmt.Println("启动服务成功...")
			os.Exit(0)
		}

	case "stop":

		err = control(svcName, svc.Stop, svc.Stopped)

		if err == nil {
			fmt.Println("停止服务成功...")
			os.Exit(0)
		}
	}

	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}

	os.Exit(0)
}
