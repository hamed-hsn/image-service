package v1

import (
	"github.com/labstack/echo/v4"
	"image_service/internal/controller/server"
	"image_service/internal/dto"
	"net/http"
)

func UploadImageApi(ctrl *server.Controller) echo.HandlerFunc {
	return func(c echo.Context) error {
		ff, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		file, err := ff.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		contentType := ff.Header.Get("Content-Type")
		resp, err := ctrl.UploadImage(c.Request().Context(), dto.UploadRequest{File: file, Size: ff.Size, ContentType: contentType})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, resp)
	}
}
