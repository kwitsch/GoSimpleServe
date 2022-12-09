//go:build linux
// +build linux

package main

import (
	"os"
	"time"
	_ "time/tzdata"

	reaper "github.com/ramr/go-reaper"
)

func init() {
	go reaper.Reap()

	setLocaltime()
}

func setLocaltime() {
	if lt, err := os.ReadFile("/etc/localtime"); err == nil {
		if t, err := time.LoadLocationFromTZData("", lt); err == nil {
			time.Local = t

			return
		}
	}

	if tz := os.Getenv("TZ"); tz != "" {
		if t, err := time.LoadLocation(tz); err == nil {
			time.Local = t

			return
		}
	}
}
