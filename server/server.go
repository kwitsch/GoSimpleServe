package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kwitsch/GoSimpleServe/files"
	"github.com/kwitsch/GoSimpleServe/util"
)

const serverPort = 80

type Server struct {
	log util.Log
	mux *http.ServeMux
}

func New(verbose bool) *Server {
	s := Server{
		log: *util.NewLog("server", verbose),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(files.StaticFilesDir)))
	mux.HandleFunc("/files", s.getFiles)

	s.mux = mux

	return &s
}

func (s *Server) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", serverPort), s.mux)
}

func (s *Server) getFiles(w http.ResponseWriter, r *http.Request) {
	s.log.V("got /files request")
	f := files.GetFiles()
	s.log.V("responese:\n", f)
	io.WriteString(w, f)
}
