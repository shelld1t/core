package httpServer

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"shelld1t/mstemplate/pkg/shelld1t/log"
	"shelld1t/mstemplate/pkg/shelld1t/middleware"
)

type Server struct {
	Router *echo.Echo
	logger *log.Logger
}

func New(logger *log.Logger) (*Server, error) {
	s := &Server{
		Router: configureEcho(),
		logger: logger,
	}
	return s, nil
}

func (s *Server) Run(addr string) error {
	s.logger.Info(fmt.Sprintf("http server starting at port %s", addr))
	err := s.Router.Start(addr)
	if err != nil {
		return errors.Wrap(err, "error run httpServer")
	}
	return nil
}

func configureEcho() *echo.Echo {
	e := echo.New()
	e.Debug = false
	e.HTTPErrorHandler = middleware.ErrorHandler()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.SetTraceId())
	e.Use(middleware.RequestLogger([]string{}))
	return e
}

func (s *Server) GroupRouter(prefix string, middleware ...echo.MiddlewareFunc) *echo.Group {
	return s.Router.Group(prefix, middleware...)
}

func (s *Server) AddRouter(method, path string, handler HandlerFunc, middleware ...echo.MiddlewareFunc) {
	s.Router.Add(method, path, wrapHandler(handler), middleware...)
}

func (s *Server) AddEndpoints(endpoints []*Endpoint) {
	for _, endpoint := range endpoints {
		s.Router.Add(endpoint.Method, endpoint.Path, wrapHandler(endpoint.Handle))
	}
}
