package service

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func GetImageByFilePath(filePath string) (string, error) {
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("FilePath:" + filePath)
		fmt.Println("There was an error getting the image")
		fmt.Println(err.Error())
		return "error", err
	}

	var base64Encoding string

	mimeType := http.DetectContentType(imageData)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(imageData)

	return base64Encoding, nil

	/*	imageResponse := base64.StdEncoding(len(imageData))

		output := make([]byte, base64.StdEncoding.EncodedLen(len(imageData)))
		base64.StdEncoding.Encode(output, imageData)

		return output, err
	*/
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
