package apiserver

import (
	"github.com/sirupsen/logrus"
)

// server
type APIServer struct {
	config *Config,
	logger *logrus.Logger,
}

// init
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// start server
func (s *APIServer) Start() error {

	return nil
}

func (s *ApiServer) configurateLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}
