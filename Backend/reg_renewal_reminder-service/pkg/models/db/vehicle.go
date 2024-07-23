package db

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Vehicle struct {
	OwnerName        string `dynamodbav:"owner_name"`
	OfficeName       string `dynamodbav:"office_name"`
	RegNo            string `dynamodbav:"reg_no"`
	VehicleClassDesc string `dynamodbav:"vehicle_class_desc"`
	Model            string `dynamodbav:"model"`
	Reg_Upto         string `dynamodbav:"reg_upto"`
	VehicleType      string `dynamodbav:"vehicle_type"`
	Mobile           string `dynamodbav:"mobile"`
	PucUpto          string `dynamodbav:"puc_upto"`
	LastCheckDate    string `dynamodbav:"last_check_date"`
}

func GetVehicleOnRegNo(reg_no string) (*Vehicle, error) {
	var vehicles []*Vehicle
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("VEHICLE_TABLE_ARN")

	// Define the key condition expression
	keyCond := expression.Key("reg_no").Equal(expression.Value(reg_no))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		log.Println("error creating key condition expression")
		return nil, err
	}

	input := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(tableName),
		ConsistentRead:            aws.Bool(true),
		Limit:                     aws.Int32(1),
	}

	output, err := svc.Query(context.TODO(), input)
	if err != nil {
		log.Println("error querying items")
		return nil, err
	}
	err = attributevalue.UnmarshalListOfMaps(output.Items, &vehicles)
	if err != nil {
		log.Println("error unmarshalling attributes : ", err)
		return nil, err
	}

	//* Check if there are no vehicles
	if len(vehicles) == 0 {
		log.Println("NO vehicles")
		return &Vehicle{}, nil
	}

	return vehicles[0], nil
}
