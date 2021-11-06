package monitoring

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"shelld1t/mstemplate/pkg/shelld1t/httpServer"
	"shelld1t/mstemplate/pkg/shelld1t/model"
)

type Health struct {
}

func newHealthController() *Health {
	return &Health{}
}

func (h *Health) HealthEndpoints() []*httpServer.Endpoint {
	return []*httpServer.Endpoint{
		{
			Path:   "/health",
			Method: http.MethodGet,
			Handle: h.Ping,
		},
	}
}

func (h *Health) Ping(ectx echo.Context) model.Response {
	return model.Ok("ok")
}
