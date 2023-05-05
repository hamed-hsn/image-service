package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"image_service/internal/controller/server"
)

func New(cntrl *server.Controller) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	registerRoutes(e, cntrl)
	return e
}
