package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getPwd() string {
	str, _ := os.Getwd()
	return str
}

func isEnabled() bool {
	cmd := exec.Command("nmcli", "radio", "wifi")
	// todo check on linux if this still applies if not check for cmd.Output()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		return false
	}

	return true
	// END_TODO

	enabledOutput, err := scriptRunner(isEnabledScript)
	if err != nil {
		log.Fatal(err)
	}
	str := string(enabledOutput)
	str = strings.TrimSuffix(str, "\n")
	return str == "enabled"
}

func enableWiFi() {
	if _, err := scriptRunner(enableScript); err != nil {
		log.Fatal(err)
	}
}

func disableWiFi() {
	if _, err := scriptRunner(disableScript); err != nil {
		log.Fatal(err)
	}
}

func isAllowed(t, from, to time.Time) bool {
	if t.After(from) && t.Before(to) {
		return true
	}

	return false
}

func isConnected() bool {
	cmd := exec.Command("wget", "--spider", "http://google.com")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		return false
	}

	return true
}
