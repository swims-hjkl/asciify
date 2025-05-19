package asciify

import (
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
