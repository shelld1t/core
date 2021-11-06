package httpServer

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"shelld1t.io/core/model"
)

type HandlerFunc = func(ctx echo.Context) model.Response

func wrapHandler(handlerFunc HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res := handlerFunc(ctx)
		enc, err := res.Encode()
		if err != nil {
			return errors.Wrap(err, "error encode response")
		}
		return ctx.Blob(res.Code(), echo.MIMEApplicationJSON, enc)
	}
}
