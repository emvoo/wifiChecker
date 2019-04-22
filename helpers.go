package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func getPwd() string {
	str, _ := os.Getwd()
	return str
}

func scriptRunner(script string) ([]byte, error) {
	pathToScript := filepath.Join(getPwd(), scriptsPath, script)
	return exec.Command(pathToScript).Output()
}

func isEnabled() bool {
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
	out, err := scriptRunner(isConnectedScript)
	if err != nil {
		log.Fatal(err)
	}

	if string(out) == online {
		return true
	}

	return false
}
