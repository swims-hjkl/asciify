package asciify

import (
	"errors"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func writePNGImage(inputImage image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return errors.New("error creating output image!")
	}
	defer file.Close()
	png.Encode(file, inputImage)
	return nil
}

func readImage(sourcePath string) (image.Image, error) {
	file, err := os.Open(sourcePath)
	if err != nil {
		return nil, errors.New("something went wrong reading the file")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, errors.New("something went wrong decoding the image")
	}
	return img, nil
}

func imageToAscii(originalImage image.Image, RW int) string {
	OW := originalImage.Bounds().Dx()
	OH := originalImage.Bounds().Dy()
	RH := int(float32(RW)*float32(OH)/float32(OW)) / 2
	resizedImage := image.NewRGBA(image.Rect(0, 0, RW, RH))
	scaleX := float32(OW) / float32(RW)
	scaleY := float32(OH) / float32(RH)
	asciiArtChars := "$@&M*oahwmO0UYcvf/(1{[]-_+~<>i!I;,`'."
	characterIndexDivideFactor := float32(256 / len(asciiArtChars))

	outputString := ""

	for rowIdx := range resizedImage.Bounds().Dy() {
		rowString := ""
		for colIdx := range resizedImage.Bounds().Dx() {
			originalX := int(float32(colIdx) * scaleX)
			originalY := int(float32(rowIdx) * scaleY)
			R, G, B, _ := originalImage.At(originalX, originalY).RGBA()
			r, g, b := float32(R>>8), float32(G>>8), float32(B>>8)
			grayPixel := (r * 0.2989) + (g * 0.5870) + (b * 0.1140)
			idx := int(grayPixel / characterIndexDivideFactor)
			if idx >= len(asciiArtChars) {
				idx = len(asciiArtChars) - 1
			}
			rowString = rowString + (string(asciiArtChars[idx]))
		}
		outputString = outputString + rowString + "\n"
	}
	return outputString
}
