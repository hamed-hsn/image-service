package rest

import (
	"github.com/labstack/echo/v4"
	"image_service/internal/controller/server"
	v1 "image_service/internal/delivery/rest/v1"
)

func registerRoutes(e *echo.Echo, ctrl *server.Controller) {
	groupV1 := e.Group("/v1")
	groupV1.GET("/list", v1.ListApi(ctrl))
	groupV1.GET("/info", v1.GetInfoApi(ctrl))
	groupV1.GET("/view", v1.GetImageApi(ctrl))
	groupV1.POST("/upload", v1.UploadImageApi(ctrl))
}
