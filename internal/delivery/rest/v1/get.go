package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"image_service/internal/controller/server"
	"image_service/internal/dto"
	"io"
	"net/http"
	"os"
)

func GetInfoApi(ctrl *server.Controller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.GetImageRequest
		request.Url = c.QueryParam("url")
		request.CommonKey = c.QueryParam("common_key")
		resp, err := ctrl.GetImageInfo(c.Request().Context(), request)
		if err != nil {
			//TODO only for development phase return error message
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, resp)
	}
}

func GetImageApi(ctrl *server.Controller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dto.GetImageRequest
		request.Url = c.QueryParam("url")
		request.CommonKey = c.QueryParam("common_key")
		resp, err := ctrl.GetImageInfo(c.Request().Context(), request)
		if err != nil {
			//TODO only for development phase return error message
			return c.JSON(http.StatusInternalServerError, err)
		}
		file, err := os.Open(resp.Info.LocalPath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.Response().Header().Set("extension", resp.Info.Ext)
		c.Response().Header().Set("mime-type", resp.Info.MimeType)
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", resp.Info.Size))
		raw, _ := io.ReadAll(file)
		return c.Blob(http.StatusOK, resp.Info.MimeType, raw)
	}
}
