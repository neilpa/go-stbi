// Package bmp provides a GIF decoder via the go bindings for stb_image.h
// and registers for use with image.Decode and image.DecodeConfig.
package gif // import "neilpa.me/go-stbi/gif"

import (
	"encoding/binary"
	"errors"
	"io"
	"image"
	"image/color"

	"neilpa.me/go-stbi"
)

// Header is the magic string at the start of a GIF file.
const Header = "GIF8?a"

// ErrInvalid is returned from DecodeConfig for non GIF files.
var ErrInvalid = errors.New("Invalid GIF")

// Decode reads a GIF image from r and returns an image.RGBA.
func Decode(r io.Reader) (image.Image, error) {
	return stbi.LoadReader(r)
}

// DecodeConfig returns the dimensions and an RGBA color model of the GIF
// backed by reader. Returns ErrInvalid if the file isn't a GIF.
func DecodeConfig(r io.Reader) (image.Config, error) {
	// TODO Make sure we get the right color model?
	cfg := image.Config{ ColorModel: color.RGBAModel }

	var h gifHeader
	err := binary.Read(r, binary.LittleEndian, &h)
	if err != nil {
		return cfg, err
	}
	if string(h.Magic[:]) != "GIF87a" && string(h.Magic[:]) != "GIF89a" {
		return cfg, ErrInvalid
	}
	cfg.Width, cfg.Height = int(h.Width), int(h.Height)
	return cfg, nil
}

func init() {
	image.RegisterFormat("gif", Header, Decode, DecodeConfig)
}

type gifHeader struct {
	Magic [6]byte
	Width uint16
	Height uint16
}
