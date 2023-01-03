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
		log: *util.NewLog("Server:", config.IsVerbose()),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(files.StaticFilesDir)))

	if !files.HasIndex() {
		s.log.M("Has no index file")
		mux.HandleFunc("/index.html", s.getFakeIndex)
	}

	if config.FilesEndpointEnabled() {
		s.log.M("Files enpoint is enabled")
		s.log.V("Files:\n" + files.GetFiles() + "\n---------")
		mux.HandleFunc("/files", s.getFiles)
	}

	if config.HasConfigTemplate() {
		s.log.M("Has config")
		s.log.V("Config:\n" + config.GetConfig() + "\n---------")
		mux.HandleFunc("/config", s.getConfig)
	}

	s.mux = mux

	return &s
}

func (s *Server) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", serverPort), s.mux)
}

func (s *Server) getFakeIndex(w http.ResponseWriter, r *http.Request) {
	s.log.V("Response for /index.html: OK")
	io.WriteString(w, "OK")
}

func (s *Server) getFiles(w http.ResponseWriter, r *http.Request) {
	f := files.GetFiles()
	s.log.V("Response for /files:\n" + f)
	io.WriteString(w, f)
}

func (s *Server) getConfig(w http.ResponseWriter, r *http.Request) {
	c := config.GetConfig()
	s.log.V("Response for /config:\n" + c)
	io.WriteString(w, c)
}
