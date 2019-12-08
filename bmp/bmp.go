package bmp

import (
	"encoding/binary"
	"errors"
	"io"
	"image"
	"image/color"

	"github.com/neilpa/go-stbi"
)

// Header is the magic string at the start of a BMP file.
const Header = "BM"

// ErrInvalid is returned from DecodeConfig for non BMP files.
var ErrInvalid = errors.New("Invalid BMP")

// Decode reads a BMP image from r and returns an image.RGBA.
func Decode(r io.Reader) (image.Image, error) {
	return stbi.LoadReader(r)
}

// DecodeConfig returns the dimensions and an RGBA color model of the BMP
// backed by reader. Returns ErrInvalid if the file isn't a BMP.
func DecodeConfig(r io.Reader) (image.Config, error) {
	cfg := image.Config{ ColorModel: color.RGBAModel }

	var h bmpHeader
	err := binary.Read(r, binary.LittleEndian, &h)
	if err != nil {
		return cfg, err
	}
	if string(h.Magic[:]) != Header {
		return cfg, ErrInvalid
	}

	// https://en.wikipedia.org/wiki/BMP_file_format#DIB_header_(bitmap_information_header)
	switch h.DIBSize {
	case 12:
		var dim struct{ X, Y uint16 }
		err = binary.Read(r, binary.LittleEndian, &dim)
		cfg.Width, cfg.Height = int(dim.X), int(dim.Y)

	case 40, 56, 108, 124:
		var dim struct{ X, Y int32 }
		err = binary.Read(r, binary.LittleEndian, &dim)
		cfg.Width, cfg.Height = int(dim.X), int(dim.Y)

	default:
		err = ErrInvalid
	}

	return cfg, err
}

func init() {
	image.RegisterFormat("bmp", Header, Decode, DecodeConfig)
}

type bmpHeader struct {
	Magic [2]byte
	FileSize uint32
	Reserved1 uint16
	Reserved2 uint16
	DataOffset uint32
	DIBSize uint32
}
