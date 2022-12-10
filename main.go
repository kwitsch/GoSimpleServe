//go:build linux
// +build linux

package main

import (
	"os"

	"github.com/kwitsch/GoSimpleServe/cmd"
	"github.com/kwitsch/GoSimpleServe/config"
	"github.com/kwitsch/GoSimpleServe/util"
	reaper "github.com/ramr/go-reaper"
)

func init() {
	go reaper.Reap()
}

func main() {
	exitcode, err := cmd.RunCmd()

	if err != nil {
		log := util.NewLog("", config.IsVerbose())
		log.E(err)
	}

	os.Exit(exitcode)
}
