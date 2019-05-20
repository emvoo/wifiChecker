package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

const (
	version        = "0.0.1"
	timeFormat     = "15:04"
	description    = "Application to run the scripts on 60 seconds (default, can be overridden) intervals."
	isEnabledCmd   = "nmcli radio wifi"
	worldClockAPI  = "http://worldtimeapi.org/api/ip"
	enableWiFiCmd  = "nmcli radio wifi on"
	disableWiFiCmd = "nmcli radio wifi off"
	isConnectedCmd = "wget --spider http://google.com"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := cli.NewApp()
	app.Description = description
	app.Version = version
	app.Flags = flags

	app.Action = cli.ActionFunc(run)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	select {}
}

func run(cliCtx *cli.Context) error {
	from, to, err := getFromTo(cliCtx)
	if err != nil {
		return err
	}

	weekend := cliCtx.Bool("weekend")
	// initialize ticker interval
	interval := toInterval(cliCtx.Int64("interval"))
	// initialize ticker
	ticker := time.NewTicker(interval)
	done := make(chan bool)
	for {
		select {
		case <-done:
			log.Println("Ticking done")
			ticker.Stop()
			return nil
		case <-ticker.C:
			t := getCurrentTime()
			// check if allowed to watch
			if isAllowed(t, from, to, weekend) {
				// check if wifi enabled
				if !isEnabled() {
					log.Println("wifi is disabled, enabling...")
					enableWiFi()
					log.Println("enabled")
				}
				continue
			}
			if isEnabled() {
				log.Println("wifi is enabled, disabling...")
				disableWiFi()
				log.Println("disabled")
			}
		}
	}
}

func getFromTo(cliCtx *cli.Context) (from time.Time, to time.Time, err error) {
	from, err = toTime(cliCtx.String("from"))
	if err != nil {
		return
	}

	to, err = toTime(cliCtx.String("to"))
	if err != nil {
		return
	}

	return
}
