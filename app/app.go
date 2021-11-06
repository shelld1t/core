package app

import (
	"github.com/pkg/errors"
	"shelld1t.io/core/httpServer"
	"shelld1t.io/core/log"
)

type App struct {
	server *httpServer.Server
	Config *Config
	Log    *log.Logger
}

type Handler struct {
	endpoint httpServer.Endpoint
}

type HandlerFunc func(a *httpServer.Server) error

type initHandlersFunc func(server *httpServer.Server) error

func New(config *Config) (*App, error) {
	l, err := configureLogger(config)
	if err != nil {
		return nil, errors.Wrap(err, "error create logger")
	}

	server, err := httpServer.New(l)
	if err != nil {
		return nil, errors.Wrap(err, "error create http server")
	}

	return &App{
		Config: config,
		Log:    l,
		server: server,
	}, nil
}

func (a *App) Run() error {
	return a.server.Run(a.Config.BindAddr)
}

func (a *App) InitHttpHandlers(f initHandlersFunc) error {
	return f(a.server)
}

func configureLogger(config *Config) (*log.Logger, error) {
	rootLog, err := log.NewLogger(config.LoggerCfg)

	if err != nil {
		return nil, err
	}
	log.SetRootLog(rootLog)
	return rootLog, nil
}
