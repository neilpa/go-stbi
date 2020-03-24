# go-stbi

[![CI](https://github.com/neilpa/go-stbi/workflows/CI/badge.svg)](https://github.com/neilpa/go-stbi/actions/)
[![GoDoc](https://godoc.org/neilpa.me/go-stbi?status.svg)](https://godoc.org/neilpa.me/go-stbi)

Go binding for [stb_image.h][].

## Usage

Load an `image.RGBA` from some path on disk.

```go
import "neilpa.me/go-stbi"

image, err := stbi.Load("path/to/image.jpeg")
// ...
```

There are also format specific sub-packages that register decoders for use
with the standard `image.Decode` and `image.DecodeConfig` methods.

```go
import (
    "image"

    _ "neilpa.me/go-stbi/bmp"
    _ "neilpa.me/go-stbi/gif"
    _ "neilpa.me/go-stbi/jpeg"
    _ "neilpa.me/go-stbi/png"
)

bmp, _, err := image.Decode("path/to/image.bmp")
gif, _, err := image.Decode("path/to/image.gif")
jpg, _, err := image.Decode("path/to/image.jpg")
png, _, err := image.Decode("path/to/image.png")
// ...
```

## Licence

This code is released into the public domain.

[stb_image.h]: https://github.com/nothings/stb/blob/f67165c2bb2af3060ecae7d20d6f731173485ad0/stb_image.h
