package asciify

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"testing"
)

func TestImageToASCII(t *testing.T) {
	basePath := "testdata"
	testImagePath := filepath.Join(basePath, "portrait_table.jpg")
	expectedOutputFilePath := filepath.Join(basePath, "expected_output.txt")

	expectedString, err := os.ReadFile(expectedOutputFilePath)
	if err != nil {
		t.Fatalf("failed to read expected output file: %v", err)
	}

	outputString, err := ConvertImageToAscii(testImagePath, 150)
	if err != nil {
		t.Fatalf("failed to convert image to ascii: %v", err)
	}

	if !(outputString == string(expectedString)) {
		t.Fatal("The expected string does not match the actual string")
	}
}

func TestImageToASCIIConcurrent(t *testing.T) {
	basePath := "testdata"
	testImagePath := filepath.Join(basePath, "portrait_table.jpg")
	expectedOutputFilePath := filepath.Join(basePath, "expected_output.txt")

	expectedString, err := os.ReadFile(expectedOutputFilePath)
	if err != nil {
		t.Fatalf("failed to read expected output file: %v", err)
	}

	outputString, err := ConvertImageToAsciiConcurrent(testImagePath, 150)
	if err != nil {
		t.Fatalf("failed to convert image to ascii: %v", err)
	}

	if !(outputString == string(expectedString)) {
		t.Fatal("The expected string does not match the actual string")
	}
}

func BenchmarkGenerateAscii(b *testing.B) {
	largeImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{7680, 4320}})
	for rowIdx := range largeImage.Bounds().Dy() {
		for colIdx := range largeImage.Bounds().Dx() {
			largeImage.SetRGBA(colIdx, rowIdx, color.RGBA{255, 255, 255, 255})
		}
	}
	for range 1000 {
		imageToAscii(largeImage, 150)
	}
}

func BenchmarkGenerateAsciiConcurrent(b *testing.B) {
	largeImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{7680, 4320}})
	for rowIdx := range largeImage.Bounds().Dy() {
		for colIdx := range largeImage.Bounds().Dx() {
			largeImage.SetRGBA(colIdx, rowIdx, color.RGBA{255, 255, 255, 255})
		}
	}
	for range 1000 {
		imageToAsciiConcurrent(largeImage, 150)
	}
}
