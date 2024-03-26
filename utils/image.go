package utils

import (
	"bytes"
	"io"

	"github.com/labstack/gommon/log"
)

func ReadFileData(r io.Reader) []byte {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, r); err != nil {
		log.Errorf("Failed to read file data: %v", err)
		return nil
	}
	return buf.Bytes()
}