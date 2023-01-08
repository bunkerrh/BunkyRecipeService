package service

import (
	"encoding/base64"
	"fmt"
	"os"
)

func GetImageByFilePath(filePath string) ([]byte, error) {
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("FilePath:" + filePath)
		fmt.Println("There was an error getting the image")
		fmt.Println(err.Error())
		return nil, err
	}

	output := make([]byte, base64.StdEncoding.EncodedLen(len(imageData)))
	base64.StdEncoding.Encode(output, imageData)

	return output, err
}
