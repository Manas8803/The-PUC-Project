package image

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Image struct {
	ImageName  string `json:"image_name"`
	ImageBytes string `json:"image_bytes"`
}

func (img *Image) FromJson(req *events.APIGatewayProxyRequest) error {
	err := json.Unmarshal([]byte(req.Body), img)
	if err != nil {
		return err
	}

	return nil
}

func (img *Image) DecodeAndSaveImage() error {
	// Decode the base64 string to []byte
	data, err := base64.StdEncoding.DecodeString(img.ImageBytes)
	if err != nil {
		log.Println("error in decoding string: ", err)
		return err
	}

	// Determine the image format
	reader := bytes.NewReader(data)
	_, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("error in determining image format", err)
		return err
	}

	// Reset the reader to read from the beginning again
	reader.Seek(0, 0)

	// Convert []byte data to an image
	decodedImage, _, err := image.Decode(reader)
	if err != nil {
		log.Println("error in decoding image", err)
		return err
	}

	// Create the file in the /tmp directory
	fileName := strings.TrimSpace(img.ImageName)
	if fileName == "" {
		fileName = "number_plate." + format // Use the detected format
	}
	
	filePath := "/tmp/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("error in creating path: ", filePath)
		return err
	}
	defer file.Close()

	// Encode the image in its original format
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(file, decodedImage, nil) // Use nil for default options
	case "png":
		err = png.Encode(file, decodedImage)
	default:
		err = fmt.Errorf("unsupported format: %s", format)
	}
	if err != nil {
		log.Println("Unsupported FORMAT:", format)
		return err
	}

	log.Println("Image saved to ", filePath)
	return nil
}
