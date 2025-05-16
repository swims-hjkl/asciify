package main

import (
	"github.com/swims-hjkl/asciify"
	"flag"
)

func main() {
	path := flag.String("path", "", "A valid path to the image")
	width := flag.Int("width", 150, "Width of the output sequence of characters (default 150)")
	flag.Parse()
	asciify.ConvertImageToAscii(*path, *width)
}
