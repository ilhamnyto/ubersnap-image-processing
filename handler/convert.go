package handler

import (
	"net/http"

	"github.com/ilhamnyto/ubersnap-image-processing/entity"
	"github.com/ilhamnyto/ubersnap-image-processing/utils"
	"github.com/labstack/echo/v4"
)

func ConvertImageHandler(c echo.Context) error {
	inputFile, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message" : "Missing input file"})
	}

	inputData, err := inputFile.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : "Error opening input file"})
	}
	defer inputData.Close()

	utils.ProcessImageTask(entity.ImageProcessingTask{
		InputData: utils.ReadFileData(inputData),
		Operation: "convert",
		Quality:   100,
		Context:   c,
	})

	resp := <-utils.ResponseChan
	if resp.Err != nil {
		return c.String(http.StatusInternalServerError, resp.Err.Error())
	}

	return c.Blob(http.StatusOK, resp.ContentType, resp.Data)
}