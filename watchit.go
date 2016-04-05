package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

const APP_NAME = "Watch It! (github.com/eduardonunesp/watchit)"
const APP_USAGE = "watchit <directory|file> <executable>"
const VERSION_MAJOR = "0"
const VERSION_MINOR = "1"
const VERSION_BUILD = "0"

func execCommand(command string) {
	cmd := command
	args := []string{}

	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		log.Fatal(err)
	}

	log.Println(string(cmdOut))
}

func runWatchit(c *cli.Context) {
	if len(c.Args()) < 2 {
		fmt.Println("usage:", c.App.UsageText)
		return
	}

	watcheable, err := os.Stat(c.Args()[0])
	if err != nil {
		fmt.Println("watchit:", err)
		return
	}

	_, err = os.Stat(c.Args()[1])
	if err != nil {
		fmt.Println("watchit:", err)
		return
	}

	watcher := &Watcher{}
	chanWatcher := watcher.startWatcher()

	done := make(chan bool)
	watcher.addWatchable(c.Args()[0], watcheable)

	for changes := range chanWatcher {
		if changes {
			execCommand(c.Args()[1])
		}
	}
	<-done
}

func main() {
	app := cli.NewApp()

	app.Name = APP_NAME
	app.UsageText = APP_USAGE
	app.Version = VERSION_MAJOR + "." + VERSION_MINOR + "." + VERSION_BUILD

	app.Action = runWatchit
	app.Run(os.Args)
}
