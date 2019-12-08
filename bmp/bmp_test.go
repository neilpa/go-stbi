package bmp_test

import (
	"image"
	"os"
	"testing"

	_ "github.com/neilpa/go-stbi/bmp"
)


var tests = []struct{
	path string
	width, height int
} {
	{ "../testdata/red.16x16.bmp", 16, 16 },
}

func TestDecode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			f, err := os.Open(tt.path)
			if err != nil {
				t.Error(err)
			}
			img, _, err := image.Decode(f)
			if err != nil {
				t.Error(err)
			}
			s := img.Bounds().Size()
			if s.X != tt.width || s.Y != tt.height {
				t.Errorf("bounds: got %dx%d want %dx%d", s.X, s.Y, tt.width, tt.height)
			}
			rgba, ok := img.(*image.RGBA)
			if !ok {
				t.Errorf("format: not RGBA")
			}
			p := tt.width * tt.height * 4
			if len(rgba.Pix) != p {
				t.Errorf("pixels: got %d want %d", len(rgba.Pix), p)
			}
		})
	}
}

func TestDecodeConfig(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			f, err := os.Open(tt.path)
			if err != nil {
				t.Error(err)
			}
			cfg, _, err := image.DecodeConfig(f)
			if err != nil {
				t.Error(err)
			}
			if cfg.Width != tt.width || cfg.Height != tt.height {
				t.Errorf("got %dx%d want %dx%d", cfg.Width, cfg.Height, tt.width, tt.height)
			}
		})
	}
}

