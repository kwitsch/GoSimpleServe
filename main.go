//go:build linux
// +build linux

package main

import (
	"os"
	"time"

	reaper "github.com/ramr/go-reaper"
)

func init() {
	go reaper.Reap()

	if lt, err := os.ReadFile("/etc/localtime"); err == nil {
		if t, err := time.LoadLocationFromTZData("", lt); err == nil {
			time.Local = t
		}
	}
}
