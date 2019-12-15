// Package png provides a PNG decoder via the go bindings for stb_image.h
// and registers for use with image.Decode and image.DecodeConfig.
package png // import "neilpa.me/go-stbi/png"

import (
	"encoding/binary"
	"errors"
	"io"
	"image"
	"image/color"

	"neilpa.me/go-stbi"
)

// Header is the magic string at the start of a PNG file.
const Header = "\x89PNG\r\n\x1a\n"

// ErrInvalid is returned from DecodeConfig for non PNG files.
var ErrInvalid = errors.New("Invalid PNG")

// Decode reads a PNG image from r and returns an image.RGBA.
func Decode(r io.Reader) (image.Image, error) {
	return stbi.LoadReader(r)
}

// DecodeConfig returns the dimensions and an RGBA color model of the PNG
// backed by reader. Note this simply wraps the stdlib png.DecodeConfig.
func DecodeConfig(r io.Reader) (image.Config, error) {
	cfg := image.Config{ColorModel: color.RGBAModel}

	var h pngHeader
	err := binary.Read(r, binary.BigEndian, &h)
	if err != nil {
		return cfg, err
	}

	// IHDR is the first chunk after the signature
	// https://en.wikipedia.org/wiki/Portable_Network_Graphics#File_header
	if string(h.Magic[:]) != Header || string(h.ChunkType[:]) != "IHDR" {
		return cfg, ErrInvalid
	}

	cfg.Width, cfg.Height = int(h.Width), int(h.Height)
	return cfg, nil
}

func init() {
	image.RegisterFormat("png", Header, Decode, DecodeConfig)
}

// pngHeader is enough to decode up to the widht/height
type pngHeader struct {
	Magic [8]byte
	ChunkSize uint32
	ChunkType [4]byte
	Width uint32
	Height uint32
}
