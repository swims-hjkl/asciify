package main

import (
	"flag"
	"fmt"

	"github.com/swims-hjkl/asciify"
)

func main() {
	path := flag.String("path", "", "A valid path to the image")
	width := flag.Int("width", 150, "Width of the output sequence of characters (default 150)")
	flag.Parse()
	asciiRepString := asciify.ConvertImageToAscii(*path, *width)
	fmt.Print(asciiRepString)
}
