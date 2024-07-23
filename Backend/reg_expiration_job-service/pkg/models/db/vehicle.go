package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
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

func GetAllVehicles() (*[]Vehicle, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("VEHICLE_TABLE_ARN")),
	}

	var vehicles []Vehicle
	paginator := dynamodb.NewScanPaginator(client, input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		for _, item := range output.Items {
			var vehicle Vehicle
			err = attributevalue.UnmarshalMap(item, &vehicle)
			if err != nil {
				return nil, err
			}
			vehicles = append(vehicles, vehicle)
		}
	}

	return &vehicles, nil
}

func UpdateLastCheckDate(vehicle Vehicle) error {
	log.Println("HEEE")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("ERROR")
		return err
	}

	client := dynamodb.NewFromConfig(cfg)

	currentDate := time.Now().Format("02-01-2006")

	log.Println("REG NO : ", vehicle.RegNo)
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("VEHICLE_TABLE_ARN")),
		Key: map[string]types.AttributeValue{
			"reg_no": &types.AttributeValueMemberS{Value: vehicle.RegNo},
		},
		UpdateExpression: aws.String("SET last_check_date = :cd"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":cd": &types.AttributeValueMemberS{Value: currentDate},
		},
	}

	_, err = client.UpdateItem(context.TODO(), input)
	if err != nil {
		log.Println("Error updating : ", err)
		return err
	}

	return nil
}
