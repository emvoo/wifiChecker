package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func isEnabled() bool {
	comm, args := toCommand(isEnabledCmd)
	b, err := exec.Command(comm, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	str = strings.TrimSuffix(str, "\n")
	return str == "enabled"
}

func enableWiFi() {
	commandRunner(enableWiFiCmd)
}

func disableWiFi() {
	commandRunner(disableWiFiCmd)
}

func isAllowed(t, from, to time.Time) bool {
	if isWeekend(t) && t.Before(to) {
		return true
	}

	if t.After(from) && t.Before(to) {
		return true
	}

	return false
}

func isWeekend(t time.Time) bool {
	// Saturday or Sunday
	if t.Weekday() == 6 || t.Weekday() == 0 {
		return true
	}
	return false
}

func isConnected() bool {
	return commandRunner(isConnectedCmd)
}

func commandRunner(command string) bool {
	comm, args := toCommand(command)

	cmd := exec.Command(comm, args...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		return false
	}

	return true
}

func toCommand(input string) (string, []string) {
	parts := strings.Split(input, " ")
	if len(parts) <= 1 {
		log.Fatal(fmt.Sprintf("command [%s] not recognized", input))
	}

	comm := parts[0]
	args := parts[1:]

	return comm, args
}
