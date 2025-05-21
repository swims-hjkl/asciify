package asciify

import "errors"

func ConvertImageToAscii(imagePath string, width int) (string, error) {
	if fileNotExists(imagePath) {
		return "", errors.New("imagePath is not valid")
	}
	if width < 10 {
		return "", errors.New("width cannot be lesser than 10")
	}
	originalImage, err := readImage(imagePath)
	if err != nil {
		return "", err
	}
	asciiRepString := imageToAscii(originalImage, width)
	return asciiRepString, nil
}
