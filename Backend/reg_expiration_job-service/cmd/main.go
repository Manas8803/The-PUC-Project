package main

import (
	"context"
	"log"
	"net/http"
	"time"

	lmd "github.com/Manas8803/The-PUC-Project__BackEnd/reg_expiration_job-service/pkg/lib/lambda"
	"github.com/Manas8803/The-PUC-Project__BackEnd/reg_expiration_job-service/pkg/lib/util"
	"github.com/Manas8803/The-PUC-Project__BackEnd/reg_expiration_job-service/pkg/models/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//* Fetch all the rows
	vehicles, err := service.GetAllVehicles()
	if err != nil {
		log.Println("Error getting vehicles from service: ", err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	//* One by one check for PUC expiry
	for i := 0; i < len(*vehicles); i++ {
		is_puc_exp, err := util.IsPucExpired(&(*vehicles)[i])
		if err != nil {
			log.Println("Error checking for PUC expiry for vehicle : ", (*vehicles)[i])
			continue
		}
		if is_puc_exp {
			err := lmd.InvokeVRCHandler((*vehicles)[i])
			if err != nil {
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: http.StatusInternalServerError}, err
			}
			return events.APIGatewayProxyResponse{Body: "Succesfully invoked VRC lambda`"}, nil
		} else {
			//* Check for warning days
			check_warning_days, err := util.CheckWarningDays(&(*vehicles)[i])
			if err != nil {
				log.Println("Error checking warning days : ", err)
			}
			if check_warning_days {
				//			TODO : INVOKE PUCExpiryWarner Lambda
				log.Println("Invoking PUCExpiryWarner Lambda(x day)")
			}
		}

		err = service.UpdateLastCheckDate((*vehicles)[i])
		if err != nil {
			log.Println(err)
		}

	}

	return events.APIGatewayProxyResponse{Body: "Cron JOB executed successfully at : " + time.Now().String()}, nil
}

func main() {
	lambda.Start(Handler)

}
