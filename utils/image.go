package utils

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/ilhamnyto/ubersnap-image-processing/entity"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/semaphore"
)

const (
	MaxWorkers = 10 
)

var ResponseChan = make(chan entity.ImageResponse, 100)

var (
	semPool  = semaphore.NewWeighted(int64(MaxWorkers))
	Wg       sync.WaitGroup
	TaskChan = make(chan entity.ImageProcessingTask, 100) 
)

func ProcessImageTask(task entity.ImageProcessingTask) {
	TaskChan <- task
}

func ImageProcessingWorker() {
	defer Wg.Done()

	for task := range TaskChan {
		if err := semPool.Acquire(context.Background(), 1); err != nil {
			log.Errorf("Failed to acquire semaphore: %v", err)
			return
		}

		ProcessImage(task)

		semPool.Release(1)
	}
}

func ProcessImage(task entity.ImageProcessingTask) {
	src, err := imaging.Decode(bytes.NewReader(task.InputData))
	if err != nil {
		log.Errorf("Failed to decode input data: %v", err)
		return
	}

	var dst *image.NRGBA
	var outputData bytes.Buffer

	switch task.Operation {
	case "convert":
		dst = imaging.Clone(src)
		err = imaging.Encode(&outputData, dst, imaging.JPEG, imaging.JPEGQuality(task.Quality))
		if err != nil {
			log.Errorf("Failed to encode image: %v", err)
			ResponseChan <- entity.ImageResponse{Context: task.Context, Err: err}
			return
		}
		ResponseChan <- entity.ImageResponse{
			Context:     task.Context,
			Data:        outputData.Bytes(),
			ContentType: "image/jpeg",
		}
	case "resize":
		dst = imaging.Resize(src, task.Width, task.Height, imaging.Lanczos)
		err = imaging.Encode(&outputData, dst, imaging.PNG, imaging.JPEGQuality(task.Quality))
		if err != nil {
			log.Errorf("Failed to encode image: %v", err)
			ResponseChan <- entity.ImageResponse{Context: task.Context, Err: err}
			return
		}
		ResponseChan <- entity.ImageResponse{
			Context:     task.Context,
			Data:        outputData.Bytes(),
			ContentType: "image/png",
		}
	case "compress":
		dst = imaging.Clone(src)
		err = imaging.Encode(&outputData, dst, imaging.PNG, imaging.JPEGQuality(task.Quality))
		if err != nil {
			log.Errorf("Failed to compress image: %v", err)
			ResponseChan <- entity.ImageResponse{Context: task.Context, Err: err}
			return
		}
		ResponseChan <- entity.ImageResponse{
			Context:     task.Context,
			Data:        outputData.Bytes(),
			ContentType: "image/png",
		}
	default:
		log.Errorf("Invalid operation: %s", task.Operation)
		ResponseChan <- entity.ImageResponse{Context: task.Context, Err: fmt.Errorf("invalid operation")}
		return
	}
}

func ReadFileData(r io.Reader) []byte {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, r); err != nil {
		log.Errorf("Failed to read file data: %v", err)
		return nil
	}
	return buf.Bytes()
}