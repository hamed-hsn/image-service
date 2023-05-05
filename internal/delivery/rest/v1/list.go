package v1

import (
	"github.com/labstack/echo/v4"
	"image_service/internal/controller/server"
	"image_service/internal/dto"
	"net/http"
)

func ListApi(ctrl *server.Controller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.ListImageRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if request.Page == 0 {
			request.Page = 1
		}
		resp, err := ctrl.ListImages(c.Request().Context(), request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if len(resp.Data) == 0 {
			return c.JSON(http.StatusNotFound, resp)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
