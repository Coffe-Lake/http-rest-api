package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

// APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New (возвращение конфигурированного APIserver) ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start функция запуска http-server/bd-connect
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// конфигурация логгера
func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Hello World!"); err != nil {
			log.Fatal(err)
		}
	}
}
