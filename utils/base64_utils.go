package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func Base64Encode(photo string) string {
	mediaPath := os.Getenv("MEDIA_PATH")
	photoPath := filepath.Join(mediaPath, photo)
	file, err := os.ReadFile(photoPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	base64Encoded := base64.StdEncoding.EncodeToString(file)
	return base64Encoded
}
