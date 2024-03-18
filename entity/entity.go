package entity

import "github.com/labstack/echo/v4"

type ImageProcessingTask struct {
	Context   echo.Context
	InputData []byte
	Operation string
	Width     int
	Height    int
	Quality   int
}

type ImageResponse struct {
	Context     echo.Context
	Data        []byte
	ContentType string
	Err         error
}