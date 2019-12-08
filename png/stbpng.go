package stbpng

import (
	"io"
	"image"
	"image/color"
	"image/png"

	"github.com/neilpa/go-stbi"
)

const pngHeader = "\x89PNG\r\n\x1a\n"

// Decode reads a PNG image from r and returns an image.RGBA.
func Decode(r io.Reader) (image.Image, error) {
	return stbi.LoadReader(r)
}

// DecodeConfig returns the dimensions and an RGBA color model of the PNG
// backed by reader. Note this simply wraps the stdlib png.DecodeConfig.
func DecodeConfig(r io.Reader) (image.Config, error) {
	c, err := png.DecodeConfig(r)
	c.ColorModel = color.RGBAModel
	return c, err
}

func init() {
	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}
