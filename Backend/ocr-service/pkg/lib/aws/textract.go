package aws

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	ig "github.com/Manas8803/The-PUC-Project__BackEnd/ocr-service/pkg/lib/image"
	"github.com/Manas8803/The-PUC-Project__BackEnd/ocr-service/pkg/lib/aws/set"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func extractLineText(output *textract.DetectDocumentTextOutput) string {
	var result string
	for _, block := range output.Blocks {
		if block.Text == nil {
			continue
		}
		text := *block.Text
		if !set.Set[text] && len(text) >= 2 {
			result += text
			if len(result) >= 10 {
				break
			}
		}
	}

	return strings.ReplaceAll(result, " ", "")
}

func DetectText(img *ig.Image) (string, error) {

	imagePath := filepath.Join("/tmp", img.ImageName)
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		log.Println("error reading image : " + err.Error())
		return "", err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")), // Replace with your AWS region
	})

	if err != nil {
		log.Println("Error in creating session")
		return "", err
	}

	client := textract.New(sess)

	input := &textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: imageBytes,
		},
	}

	output, err := client.DetectDocumentText(input)
	if err != nil {
		return "", err
	}
	log.Println(output)
	lineText := extractLineText(output)
	log.Println("\n -------------> ", lineText)
	return lineText, nil
}
