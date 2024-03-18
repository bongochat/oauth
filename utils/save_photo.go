package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/bongochat/utils/resterrors"
)

func SaveUploadedFile(filePath string, userId int64) (string, resterrors.RestError) {
	photo, err := os.Open(filePath)
	if err != nil {
		return "", resterrors.NewInternalServerError("Could not read uploaded file", "", err)
	}
	defer photo.Close()

	// Decode the image
	img, _, err := image.Decode(photo)
	if err != nil {
		return "", resterrors.NewInternalServerError("Could not decoding uploaded file", "", err)
	}

	fileName := filepath.Base(filePath)
	log.Println(photo)
	log.Println(filePath, fileName)
	mediaPath := fmt.Sprintf("media/profile_pictures/%d/", userId)
	outputFilePath := filepath.Join(mediaPath, fileName)

	// Create the output directory if it doesn't exist
	outputDir := filepath.Dir(mediaPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", resterrors.NewInternalServerError("Error creating output directory", "", err)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return "", resterrors.NewInternalServerError("Could not create uploaded file", "", err)
	}
	defer outputFile.Close()

	// Save the image with the same quality
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return "", resterrors.NewInternalServerError("Could not save uploaded file", "", err)
	}

	return outputFilePath, nil
}
