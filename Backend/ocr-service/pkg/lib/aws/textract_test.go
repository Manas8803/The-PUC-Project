package aws

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Manas8803/The-PUC-Project__BackEnd/ocr-service/pkg/lib/image"
)

func TestDetectText(t *testing.T) {

	imagePath := filepath.Join("/tmp", "imageobject0_0.jpg")
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		t.Fatalf("Failed to read image file: %v", err)
	}
	log.Println(string(imageBytes))
	img := &image.Image{
		ImageBytes: string(imageBytes),
	}

	text, err := DetectText(img)
	if err != nil {
		t.Errorf("DetectText returned an error: %v", err)
	}

	if text == "" {
		t.Errorf("Expected non-empty text but got empty string")
	}

	expectedPattern := "MH31EH3056"
	if text != expectedPattern {
		t.Errorf("Expected text to match pattern '%s', but got '%s'", expectedPattern, text)
	}
}
