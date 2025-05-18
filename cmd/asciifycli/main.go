package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/swims-hjkl/asciify"
)

func main() {
	path := flag.String("path", "", "A valid path to the image")
	width := flag.Int("width", 150, "Width of the output sequence of characters (default 150)")
	flag.Parse()
	asciiRepString, err := asciify.ConvertImageToAscii(*path, *width)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		flag.Usage()
		os.Exit(1)
	}
	fmt.Print(asciiRepString)
}
