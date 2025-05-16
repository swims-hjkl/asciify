package asciify

import (
	"testing"
)

func TestImageToASCII(t *testing.T) {
	ConvertImageToAscii("testdata/arcreactor.png", 150)
}
