package imageutil

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"

	"purekit-backend/constant"
	"purekit-backend/errors"

	"github.com/deepteams/webp"
)

func DecodeImage(input io.Reader) (image.Image, string, error) {
	src, format, err := image.Decode(input)
	if err != nil {
		return nil, "", errors.NewImageProcessingError("Failed to decode image", err)
	}
	return src, format, nil
}

func EncodeImage(img image.Image, format string, quality int) ([]byte, error) {
	if quality <= 0 || quality > 100 {
		quality = constant.DefaultImageQuality
	}

	var buf bytes.Buffer
	switch format {
	case constant.ImageFormatJPEG:
		err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return nil, errors.NewImageProcessingError("Failed to encode JPEG", err)
		}
	case constant.ImageFormatPNG:
		err := png.Encode(&buf, img)
		if err != nil {
			return nil, errors.NewImageProcessingError("Failed to encode PNG", err)
		}
	case constant.ImageFormatGIF:
		err := gif.Encode(&buf, img, nil)
		if err != nil {
			return nil, errors.NewImageProcessingError("Failed to encode GIF", err)
		}
	case constant.ImageFormatWebP:
		err := webp.Encode(&buf, img, &webp.EncoderOptions{
			Quality: float32(quality),
		})
		if err != nil {
			return nil, errors.NewImageProcessingError("Failed to encode WebP", err)
		}
	case constant.ImageFormatBMP:
		err := bmp.Encode(&buf, img)
		if err != nil {
			return nil, errors.NewImageProcessingError("Failed to encode BMP", err)
		}
	default:
		return nil, errors.NewUnsupportedFormatError("Unsupported image format")
	}

	return buf.Bytes(), nil
}

func IsSupportedFormat(format string) bool {
	for _, f := range constant.SupportedImageFormats {
		if f == format {
			return true
		}
	}
	return false
}
