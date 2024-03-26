package handler

import (
	"bytes"
	"image"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/ilhamnyto/ubersnap-image-processing/utils"
	"github.com/labstack/echo/v4"
)

func ResizeImageHandler(c echo.Context) error {
	inputFile, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message" : "Missing input file"})
	}

	width := c.QueryParam("width")
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "wrong width input"})
	}

	height := c.QueryParam("height")
	heightInt, err := strconv.Atoi(height)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "wrong height input"})
	}

	inputData, err := inputFile.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : "Error opening input file"})
	}
	defer inputData.Close()

	src, err := imaging.Decode(bytes.NewReader(utils.ReadFileData(inputData)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : "Failed to decode input data"})
	}

	var dst *image.NRGBA
	var outputData bytes.Buffer

	dst = imaging.Resize(src, widthInt, heightInt, imaging.Lanczos)
	err = imaging.Encode(&outputData, dst, imaging.PNG)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : "Failed to encode image"})
	}

	return c.Blob(http.StatusOK, "image/png", outputData.Bytes())
}