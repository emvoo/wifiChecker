package main

import (
	"time"

	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.Int64Flag{
		Name:  "interval,i",
		Usage: "use to set how often the app should check wifi connection, default every 60s",
		Value: 60,
	},
	cli.StringFlag{
		Name:  "from,f",
		Usage: "what time wifi should be available at",
		Value: "08:00",
	},
	cli.StringFlag{
		Name:  "to,t",
		Usage: "what time should the wifi be disabled",
		Value: "19:00",
	},
}

func toInterval(input int64) time.Duration {
	if input <= 0 {
		return 1 * time.Minute
	}

	return time.Duration(input) * time.Second
}

func toTime(input string) (time.Time, error) {
	return time.Parse(timeFormat, input)
}
