# Image Processing API

This is a Go backend service that provides an API for image processing tasks, including converting images to JPEG format, resizing images, and compressing images.

## Features

- Convert PNG images to JPEG format
- Resize images to specified dimensions
- Compress images with adjustable quality

## Prerequisites

- Go (version 1.16 or later)
- [Echo web framework](https://echo.labstack.com/)
- [Imaging library](https://github.com/disintegration/imaging)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ilhamnyto/ubersnap-image-processing.git
```

2. Navigate to the project directory:

```bash
cd ubersnap-image-processing
```

3. Install the required dependencies:

```go
go mod tidy
```

## Usage

1. Start the server:

```go
go run main.go
```

The server will start running on `http://localhost:8000`.

2. API Endpoints:

- **GET /** - Healthcheck endpoint
- **POST /convert** - Convert a PNG image to JPEG format
- **POST /resize** - Resize an image to specified dimensions
- **POST /compress** - Compress an image with adjustable quality

### Convert Image

```bash
POST /convert
```

**Request Body:**

- `file` (multipart/form-data): The PNG image file to be converted.

**Example:**

```bash
curl -X POST http://localhost:8000/convert -F "file=@/path/to/image.png" > output.jpeg
```

### Resize Image


**Request Body:**

- `file` (multipart/form-data): The image file to be resized.
- `width` (query parameter): The desired width for the resized image.
- `height` (query parameter): The desired height for the resized image.

**Example:**

```bash
curl -X POST http://localhost:8000/resize?width=800&height=600 -F "file=@/path/to/image.png" > output.png
```

### Compress Image

```bash
POST /compress
```
**Request Body:**

- `file` (multipart/form-data): The image file to be compressed.
- `quality` (query parameter): The desired quality for the compressed image (0-100).

**Example:**

```bash
curl -X POST http://localhost:8000/compress?quality=80 -F "file=@/path/to/image.png" > output.png
```

## Implementation Details

The API is implemented using the Echo web framework and the Imaging library for image processing tasks. The server spawns a pool of worker goroutines to handle concurrent image processing requests. The processed images are returned in the response body as JPEG data.

The code utilizes a semaphore-based approach to limit the number of concurrent workers and prevent resource exhaustion. Tasks are enqueued in a buffered channel and picked up by available workers.

## Testing

To run the test suite, navigate to the project directory and execute the following command:

```go
go test
```

The test suite includes test cases for various scenarios, such as invalid input parameters and successful image processing operations.

