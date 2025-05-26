# ASCIIFY

**A library / command line interface tool for converting an image into ASCII.**

### Usage:

#### Using the library:

```go
import "github.com/swims-hjkl/asciify"

func main() {
    asciiOutput, err := asciify.ConvertImageToAscii("testImage.png", 150, false)
}
```

#### Building and using the CLI:

```bash
make build
./bin/asciifycli -path testImage.jpg -width 150
```

### Contributions

Contributions are welcome! 
1. Open a pull request to fix a bug
2. Open an issue to discuss a new feature or change.

### License

MIT License
