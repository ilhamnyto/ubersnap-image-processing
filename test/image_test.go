package main_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ilhamnyto/ubersnap-image-processing/handler"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCompressImage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/compress?quality=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	file, err := os.Open("../logo.png") 
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "logo.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	req.Body = io.NopCloser(body)

	if assert.NoError(t, handler.CompressImageHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, "image/png", rec.Header().Get("Content-Type"))
	}

}
func TestResizeImage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/resize?width=50&height=50", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	file, err := os.Open("../logo.png") 
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "logo.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	req.Body = io.NopCloser(body)

	if assert.NoError(t, handler.ResizeImageHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, "image/png", rec.Header().Get("Content-Type"))
	}

}

func TestConvertImage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/convert", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	file, err := os.Open("../logo.png") 
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "logo.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	req.Body = io.NopCloser(body)

	if assert.NoError(t, handler.ConvertImageHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, "image/jpeg", rec.Header().Get("Content-Type"))
	}

}

