package godeamon

import (
	"fmt"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

type DeamonApp struct {
	deamon *bool
}

var App *DeamonApp

func init() {
	goDaemon := flag.BoolP("deamon", "d", false, "run app as a daemon with -d.")
	App = &DeamonApp{
		deamon: goDaemon,
	}
}

func (app *DeamonApp) Run() *DeamonApp {
	flag.Parse()
	if *app.deamon {
		cmd := exec.Command(os.Args[0], flag.Args()...)
		if err := cmd.Start(); err != nil {
			fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
		os.Exit(0)
	}
	return app
}