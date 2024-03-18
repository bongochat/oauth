package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func ConvertToWebp(filename string) string {
	// Run webp command-line tool to convert the image
	outputFilename := fmt.Sprintf("%s.webp", filename)
	cmd := exec.Command("webp", "-o", outputFilename, filename)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return filename
	}
	return outputFilename
}
