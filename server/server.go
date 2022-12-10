package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kwitsch/GoSimpleServe/config"
	"github.com/kwitsch/GoSimpleServe/files"
	"github.com/kwitsch/GoSimpleServe/util"
)

const serverPort = 80

type Server struct {
	log util.Log
	mux *http.ServeMux
}

func New() *Server {
	s := Server{
		log: *util.NewLog("server", config.IsVerbose()),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(files.StaticFilesDir)))

	if config.FilesEndpointEnabled() {
		s.log.V("files enpoint is enabled")
		mux.HandleFunc("/files", s.getFiles)
	}

	if config.HasConfigTemplate() {
		s.log.V("has config")
		mux.HandleFunc("/config", s.getConfig)
	}

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

func (s *Server) getConfig(w http.ResponseWriter, r *http.Request) {
	s.log.V("got /config request")
	c := config.GetConfig()
	s.log.V("responese:\n", c)
	io.WriteString(w, c)
}
