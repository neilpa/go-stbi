package stbi_test

import (
	"fmt"
	"image"
	"os"
	"testing"

	_ "neilpa.me/go-stbi/bmp"
	_ "neilpa.me/go-stbi/jpeg"
	_ "neilpa.me/go-stbi/png"
)


var tests = []struct{
	path string
	width, height int
} {
	{ "testdata/red.16x8.bmp", 16, 8 },
	{ "testdata/red.16x8.jpg", 16, 8 },
	{ "testdata/red.16x8.png", 16, 8 },
}

func TestDecode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			fmt.Println("testing", tt.path)
			f, err := os.Open(tt.path)
			if err != nil {
				t.Fatal(err)
			}
			img, _, err := image.Decode(f)
			if err != nil {
				t.Fatal(err)
			}
			s := img.Bounds().Size()
			if s.X != tt.width || s.Y != tt.height {
				t.Errorf("bounds: got %dx%d want %dx%d", s.X, s.Y, tt.width, tt.height)
			}
			rgba, ok := img.(*image.RGBA)
			if !ok {
				t.Fatalf("format: not RGBA, %T", img)
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
