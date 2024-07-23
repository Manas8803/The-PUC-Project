package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Manas8803/The-PUC-Project__BackEnd/ocr-service/pkg/lib/aws"
	"github.com/Manas8803/The-PUC-Project__BackEnd/ocr-service/pkg/lib/image"
	"github.com/Manas8803/The-PUC-Project__BackEnd/ocr-service/pkg/lib/lambda"

	"github.com/aws/aws-lambda-go/events"
	lmd "github.com/aws/aws-lambda-go/lambda"
)

func Handler(cont context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//* Extracting JSON
	var img image.Image
	err := img.FromJson(&req)
	if err != nil {
		log.Println("Internal Server Error : Error in unmarshalling request body -> ", err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	//* Decoding and saving image in results directory
	err = img.DecodeAndSaveImage()
	if err != nil {
		log.Println("Internal Server Error : Error in decoding and saving image -> ", err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	//* Ocr api call
	vehicle_registration_number, err := aws.DetectText(&img)
	if err != nil {
		log.Println("Error in Detecting Text : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}
	log.Println("Retrieved number plate : ", vehicle_registration_number)

	//* Invoking check-puc lambda function :
	err = lambda.InvokeRegRenewalHandler(vehicle_registration_number)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{Body: "Succesfully invoked OCR API, Registration Number : " + vehicle_registration_number}, nil
}
func main() {
	lmd.Start(Handler)
}
