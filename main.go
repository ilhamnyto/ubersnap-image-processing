package main

import (
	"net/http"

	"github.com/ilhamnyto/ubersnap-image-processing/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
	})

	e.POST("/convert", handler.ConvertImageHandler)
	e.POST("/resize", handler.ResizeImageHandler)
	e.POST("/compress", handler.CompressImageHandler)

	e.Logger.Fatal(e.Start(":8000"))

}
