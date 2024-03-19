package test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/ilhamnyto/ubersnap-image-processing/handler"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func TestConvertImageHandler(t *testing.T) {
// 	body, writer := createMultipartBody("../logo.png")
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/convert", body)
// 	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	err := handler.ConvertImageHandler(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Contains(t, rec.Body.String(), "image/jpeg")
// }

// func TestResizeImageHandler(t *testing.T) {
// 	body, writer := createMultipartBody("../logo.png")
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/resize?width=200&height=300", body)
// 	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	err := handler.ResizeImageHandler(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Contains(t, rec.Body.String(), "image/jpeg")
// }

// func TestCompressImageHandler(t *testing.T) {
// 	body, writer := createMultipartBody("../logo.png")
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/compress?quality=80", body)
// 	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	err := handler.CompressImageHandler(c)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Contains(t, rec.Body.String(), "image/jpeg")
// }

func TestResizeImageHandlerWithInvalidWidth(t *testing.T) {
	body, writer := createMultipartBody("../logo.png")
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/resize?width=invalid&height=300", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.ResizeImageHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "wrong width input", response["message"])
}

func TestCompressImageHandlerWithInvalidQuality(t *testing.T) {
	body, writer := createMultipartBody("../logo.png")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/compress?quality=invalid", body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CompressImageHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.Equal(t, "wrong height input", response["message"])
}

func createMultipartBody(fileName string) (*bytes.Buffer, multipart.Writer) {
    var buf bytes.Buffer
    writer := multipart.NewWriter(&buf)
    
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    fileWriter, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
    if err != nil {
        panic(err)
    }
    
    _, err = io.Copy(fileWriter, file)
    if err != nil {
        panic(err)
    }
    
    writer.Close()

	return &buf, *writer
}