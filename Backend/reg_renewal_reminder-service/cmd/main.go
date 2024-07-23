package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	lmd "github.com/Manas8803/The-PUC-Project__BackEnd/reg_renewal_reminder-service/pkg/lib/lambda"
	"github.com/Manas8803/The-PUC-Project__BackEnd/reg_renewal_reminder-service/pkg/lib/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	RegNo string `json:"vehicleRegistrationNumber"`
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var data Payload
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		log.Println("Error in unmarshalling data : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	//* Check DB vehicle number exists
	vehicle_exists, vehicle, err := util.CheckRegNoIfExists(data.RegNo)
	if err != nil {
		log.Println("Error in checking the vehicle number : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	if !vehicle_exists {
		log.Printf("Vehicle with reg_no = %v does not exists in DB: ", data.RegNo)
		err := lmd.InvokeVRCHandler(*vehicle)
		if err != nil {
			log.Println("Error in invoking VRC Handler: ", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
		}
		return events.APIGatewayProxyResponse{Body: "Successfully invoked VRC Lambda", StatusCode: http.StatusOK}, nil
	}

	//* Check PUC Expiration Date <= Today
	is_puc_exp, err := util.IsPucExpired(vehicle)
	if err != nil {
		log.Println("Error in checking puc exp : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	if is_puc_exp {
		log.Println("PUC Expired", vehicle.PucUpto, vehicle.LastCheckDate)
		is_next_check_day_today, err := util.IsNextCheckDateToday(vehicle)
		if err != nil {
			log.Println("Error in checking is_next_check_day_today : ", err)
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
		}

		//* Check next renewal date == today
		if is_next_check_day_today {
			log.Println("Invoking VRC Lambda..")
			//* TRIGGER VRC SERVICE
			return events.APIGatewayProxyResponse{Body: "Successfully invoked VRC Lambda", StatusCode: http.StatusOK}, nil
		}
		//* TRIGGER PUCExpiryWarner SERVICE(one day)
		log.Println("Invoking PUCExpiryWarner Service(1 day)")
		return events.APIGatewayProxyResponse{Body: "Successfully invoked PUCExpiryWarner Lambda", StatusCode: http.StatusInternalServerError}, nil
	}

	//* Check warning days
	check_warning_days, err := util.CheckWarningDays(vehicle)
	if err != nil {
		log.Println("Error checking warning days : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}
	//* TRIGGER PUCExpiryWarner SERVICE
	if check_warning_days {
		log.Println("Invoking PUCExpiryWarner Lambda(x day)")
		// TODO : INVOKE PUCExpiryWarner Lambda
		return events.APIGatewayProxyResponse{Body: "Successfully invoked PUCExpiryWarner Lambda"}, nil
	}
	return events.APIGatewayProxyResponse{Body: "PUC is not expired, do nothing"}, nil
}
func main() {
	lambda.Start(Handler)
}
