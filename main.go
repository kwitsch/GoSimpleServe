//go:build linux
// +build linux

package main

import (
	"os"

	"github.com/kwitsch/GoSimpleServe/config"
	"github.com/kwitsch/GoSimpleServe/files"
	"github.com/kwitsch/GoSimpleServe/server"
	"github.com/kwitsch/GoSimpleServe/util"
	reaper "github.com/ramr/go-reaper"
)

func init() {
	go reaper.Reap()
}

func main() {
	log := util.NewLog("", config.IsVerbose())

	if !files.HasIndex() {
		log.E("No index.html found")
		os.Exit(1)
	}

	serv := server.New()

	if err := serv.Start(); err != nil {
		log.E(err)
		os.Exit(2)
	}
}
