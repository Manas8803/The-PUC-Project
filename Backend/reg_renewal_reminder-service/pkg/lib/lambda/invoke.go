package lambda

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type Payload struct {
	Body string `json:"body"`
}

type Data struct {
	VehicleRegistrationNumber string `json:"vehicleRegistrationNumber"`
}

func InvokeVRCHandler(regNo string) error {

	vehicle, err := json.Marshal(Data{VehicleRegistrationNumber: regNo})
	if err != nil {
		log.Println("error in marshalling data vehicle registration number : ", err)
		return err
	}

	payload, err := json.Marshal(Payload{Body: string(vehicle)})
	if err != nil {
		log.Println("error in marshalling data payload data : ", err)
		return err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION")),
	})

	if err != nil {
		return err
	}

	svc := lambda.New(sess)

	input := &lambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("VRC_HANDLER_ARN")),
		Payload:        payload,
		InvocationType: aws.String("Event"),
	}

	result, err := svc.Invoke(input)
	if err != nil {
		return err
	}

	log.Println("Invocation Result : ", result)

	return nil
}
