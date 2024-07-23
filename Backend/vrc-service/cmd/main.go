package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/lib/api"
	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/lib/socket"
	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/lib/util"
	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/models/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	RegNo string `json:"vehicleRegistrationNumber"`
}

type WebSocketMessage struct {
	Data string `json:"data"`
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var data Payload
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		log.Println("Error in unmarshalling data : ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, nil
	}

	//* Make API call
	vehicle, err := api.GetVehicleInfoByRegNo(data.RegNo)
	if err != nil {
		log.Println("Error getting vehicle info via API: ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	//* Update last check date
	vehicle.LastCheckDate = util.UpdateLastCheckDate()
	var wg sync.WaitGroup
	wg.Add(1)
	var save_db_err error
	go func() {
		defer wg.Done()
		save_db_err = service.SaveOrUpdateVehicle(*vehicle)
		if save_db_err != nil {
			log.Println("Error in saving vehicle: ", save_db_err)
		}
	}()

	//* Check for PUC Expiry
	is_puc_exp, err := util.IsPucExpired(vehicle)
	if err != nil {
		log.Println("Error in checking puc exp : ", err)
		wg.Wait()
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}
	if is_puc_exp {
		//	//TODO :IF YES, Trigger WEBSOCKET API
		log.Println("Invoking WEBSOCKET API...")
		reportErr := socket.ReportAuthority(vehicle)
		if reportErr != nil {
			return events.APIGatewayProxyResponse{Body: reportErr.Error(), StatusCode: http.StatusInternalServerError}, err
		}
		wg.Wait()
		return events.APIGatewayProxyResponse{Body: "Successfully invoked WEBSOCKET API", StatusCode: http.StatusOK}, nil
	}

	//* Check warning days
	check_warning_days, err := util.CheckWarningDays(vehicle)
	if err != nil {
		log.Println("Error checking warning days : ", err)
		wg.Wait()
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}
	//* TRIGGER PUCExpiryWarner SERVICE
	if check_warning_days {
		log.Println("Invoking PUCExpiryWarner Lambda(x day)")
		// TODO : INVOKE PUCExpiryWarner Lambda
		wg.Wait()
		return events.APIGatewayProxyResponse{Body: "Successfully invoked PUCExpiryWarner Lambda", StatusCode: http.StatusAccepted}, nil
	}
	if save_db_err != nil {
		return events.APIGatewayProxyResponse{Body: save_db_err.Error(), StatusCode: http.StatusInternalServerError}, save_db_err
	}
	wg.Wait()
	return events.APIGatewayProxyResponse{Body: "PUC Not expired save in DB", StatusCode: http.StatusAccepted}, nil
}
func main() {
	lambda.Start(Handler)
}
