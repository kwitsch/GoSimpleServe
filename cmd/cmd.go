package cmd

import (
	"errors"
	"os"

	"github.com/kwitsch/GoSimpleServe/files"
	"github.com/kwitsch/GoSimpleServe/server"
)

func RunCmd() (int, error) {
	switch os.Args[1] {
	case "serve":
		return serve()
	default:
		return -1, errors.New("no command set")
	}
}

func serve() (int, error) {
	if !files.HasIndex() {
		return 1, errors.New("No index.html found")
	}

	serv := server.New()

	if err := serv.Start(); err != nil {
		return 2, err
	}

	return 0, nil
}
