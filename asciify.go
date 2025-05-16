package asciify

func ConvertImageToAscii(imagePath string, width int) {
	if fileNotExists(imagePath) {
		panic("imagePath is not valid")
	}
	if width < 10 {
		panic("width cannot be lesser than 10")
	}
	originalImage := readImage(imagePath)
	resizedImage := resizeImage(originalImage, width)
	greyscaledImage := imageToGrayScale(resizedImage)
	grayscaleImageToAscii(greyscaledImage)
}
