package v1

import (
	"github.com/labstack/echo/v4"
	"image_service/internal/controller/server"
)

func ListApi(ctrl *server.Controller) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
