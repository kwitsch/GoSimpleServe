//go:build linux
// +build linux

package main

import (
	"os"
	"time"

	"github.com/kwitsch/GoSimpleServe/config"
	"github.com/kwitsch/GoSimpleServe/files"
	"github.com/kwitsch/GoSimpleServe/server"
	"github.com/kwitsch/GoSimpleServe/util"
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

func main() {
	verbose := config.IsVerbose()
	log := util.NewLog("", verbose)

	if !files.HasIndex() {
		log.E("No index.html found")
		os.Exit(1)
	}

	serv := server.New(verbose)

	if err := serv.Start(); err != nil {
		log.E(err)
		os.Exit(2)
	}
}
