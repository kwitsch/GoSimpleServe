package cmd

import (
	"errors"
	"net/http"
	"os"

	"github.com/kwitsch/GoSimpleServe/server"
	"github.com/kwitsch/GoSimpleServe/util"
)

func RunCmd() (int, error) {
	switch os.Args[1] {
	case "serve":
		return serve()
	case "healthcheck":
		return healthcheck()
	default:
		return -1, errors.New("no command set")
	}
}

func serve() (int, error) {
	serv := server.New()

	if err := serv.Start(); err != nil {
		return 2, err
	}

	return 0, nil
}

func healthcheck() (int, error) {
	log := util.NewLog("Healthcheck:", false)

	if _, err := http.Get("http://127.0.0.1/index.html"); err != nil {
		log.M("Fail")

		return 1, nil
	}

	log.M("Success")

	return 0, nil
}
