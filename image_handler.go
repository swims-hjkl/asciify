package asciify

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

func imageToGrayScale(inputImage image.Image) image.Image {
	newImage := image.NewGray(inputImage.Bounds())
	for rowIdx := range inputImage.Bounds().Dy() {
		for colIdx := range inputImage.Bounds().Dx() {
			R, G, B, _ := inputImage.At(colIdx, rowIdx).RGBA()
			r, g, b := float32(R/257), float32(G/257), float32(B/257)
			greyPixel := uint8((r * 0.2989) + (g * 0.5870) + (b * 0.1140))
			newImage.SetGray(colIdx, rowIdx, color.Gray{greyPixel})
		}
	}
	return newImage
}

func writePNGImage(inputImage image.Image, outputPath string) {
	file, err := os.Create(outputPath)
	if err != nil {
		panic("error creating output image!")
	}
	defer file.Close()
	png.Encode(file, inputImage)
}

func readImage(sourcePath string) image.Image {
	file, err := os.Open(sourcePath)
	if err != nil {
		panic("something went wrong reading the file")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic("something went wrong decoding the image")
	}
	return img
}

func resizeImage(originalImage image.Image, RW int) image.Image {
	OW := originalImage.Bounds().Dx()
	OH := originalImage.Bounds().Dy()
	RH := int(float64(RW)*float64(OH)/float64(OW)) / 2
	resizedImage := image.NewRGBA(image.Rect(0, 0, RW, RH))
	scaleX := float64(OW) / float64(RW)
	scaleY := float64(OH) / float64(RH)
	for rowIdx := range resizedImage.Bounds().Dy() {
		for colIdx := range resizedImage.Bounds().Dx() {
			originalX := int(float64(colIdx) * scaleX)
			originalY := int(float64(rowIdx) * scaleY)
			R, G, B, A := originalImage.At(originalX, originalY).RGBA()
			resizedImage.SetRGBA(colIdx, rowIdx, color.RGBA{R: uint8(R >> 8), G: uint8(G >> 8), B: uint8(B >> 8), A: uint8(A >> 8)})
		}
	}
	return resizedImage
}

func grayscaleImageToAscii(originalImage image.Image) {

	ascii_art_chars := "$@&M*oahwmO0UYcvf/(1{[]-_+~<>i!I;,`'."
	character_index_divide_factor := float64(256 / len(ascii_art_chars))

	for rowIdx := 0; rowIdx < originalImage.Bounds().Dy(); rowIdx++ {
		for colIdx := 0; colIdx < originalImage.Bounds().Dx(); colIdx++ {
			R, _, _, _ := originalImage.At(colIdx, rowIdx).RGBA()
			R8 := uint8(R >> 8)
			idx := int(float64(R8) / character_index_divide_factor)
			if idx >= len(ascii_art_chars) {
				idx = len(ascii_art_chars) - 1
			}
			fmt.Print(string(ascii_art_chars[idx]))
		}
		fmt.Println("")
	}
}
