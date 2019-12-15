package stbi // import "neilpa.me/go-stbi"

import (
	"errors"
	"image"
	"io"
	"io/ioutil"
	"os"
	"unsafe"
)

// #cgo LDFLAGS: -lm
// #define STB_IMAGE_IMPLEMENTATION
// #define STBI_FAILURE_USERMSG
// #include "stb_image.h"
import "C"

// Load wraps stbi_load to decode an image into an RGBA pixel struct.
func Load(path string) (*image.RGBA, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var x, y C.int
	data := C.stbi_load(cpath, &x, &y, nil, 4)
	if data == nil {
		msg := C.GoString(C.stbi_failure_reason())
		return nil, errors.New(msg)
	}
	defer C.stbi_image_free(unsafe.Pointer(data))

	return &image.RGBA{
		Pix:	C.GoBytes(unsafe.Pointer(data), y*x*4),
		Stride: 4,
		Rect:   image.Rect(0, 0, int(x), int(y)),
	}, nil
}

// LoadFile wraps stbi_load_from_file to decode an image into an RGBA pixel
// struct.
func LoadFile(f *os.File) (*image.RGBA, error) {
	mode := C.CString("rb")
	defer C.free(unsafe.Pointer(mode))
	fp, err := C.fdopen(C.int(f.Fd()), mode)
	if err != nil {
		return nil, err
	}

	var x, y C.int
	data := C.stbi_load_from_file(fp, &x, &y, nil, 4)
	if data == nil {
		msg := C.GoString(C.stbi_failure_reason())
		return nil, errors.New(msg)
	}
	defer C.stbi_image_free(unsafe.Pointer(data))

	return &image.RGBA{
		Pix:	C.GoBytes(unsafe.Pointer(data), y*x*4),
		Stride: 4,
		Rect:   image.Rect(0, 0, int(x), int(y)),
	}, nil
}

// LoadMemory wraps stbi_load_from_memory to decode an image into an RGBA
// pixel struct.
func LoadMemory(b []byte) (*image.RGBA, error) {
	var x, y C.int
	mem := (*C.uchar)(unsafe.Pointer(&b[0]))
	data := C.stbi_load_from_memory(mem, C.int(len(b)), &x, &y, nil, 4)
	if data == nil {
		msg := C.GoString(C.stbi_failure_reason())
		return nil, errors.New(msg)
	}
	defer C.stbi_image_free(unsafe.Pointer(data))

	return &image.RGBA{
		Pix:	C.GoBytes(unsafe.Pointer(data), y*x*4),
		Stride: 4,
		Rect:   image.Rect(0, 0, int(x), int(y)),
	}, nil
}

// LoadReader delegates to LoadFile if r is an *os.File, otherwise,
// LoadMemory after reading the contents.
func LoadReader(r io.Reader) (*image.RGBA, error) {
	if f, ok := r.(*os.File); ok {
		return LoadFile(f)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return LoadMemory(b)
}
