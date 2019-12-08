# go-stbi

Go binding for [stb_image.h][].

## Usage

Load an `image.RGBA` from some path on disk.

```go
import "github.com/neilpa/go-stbi"

image, err := stbi.Load("path/to/image.jpeg")
// ...
```

There are also format specific sub-packages that register decoders for use
with the standard `image.Decode` and `image.DecodeConfig` methods.

* bmp
* jpeg
* png


## Licence

This code is released into the public domain.

[stb_image.h]: https://github.com/nothings/stb/blob/f67165c2bb2af3060ecae7d20d6f731173485ad0/stb_image.h

