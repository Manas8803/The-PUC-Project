package db

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Vehicle struct {
	OwnerName        string `dynamodbav:"owner_name"`
	OfficeName       string `dynamodbav:"office_name"`
	RegNo            string `dynamodbav:"reg_no"`
	VehicleClassDesc string `dynamodbav:"vehicle_class_desc"`
	Model            string `dynamodbav:"model"`
	RegUpto          string `dynamodbav:"reg_upto"`
	VehicleType      string `dynamodbav:"vehicle_type"`
	Mobile           string `dynamodbav:"mobile"`
	PucUpto          string `dynamodbav:"puc_upto"`
	LastCheckDate    string `dynamodbav:"last_check_date"`
}

func FetchVehicles(officeName string) ([]*Vehicle, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	// Build the expression to filter on office_name (case-insensitive)
	filt := expression.Name("office_name").Equal(expression.Value(strings.ToLower(officeName)))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		return nil, err
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(os.Getenv("VEHICLE_TABLE_ARN")),
	}

	// Make the DynamoDB Scan call
	result, err := svc.Scan(params)
	if err != nil {
		return nil, err
	}

	vehicles := make([]*Vehicle, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &vehicles)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
