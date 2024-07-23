package db

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

func SaveOrUpdateVehicle(vehicle Vehicle) error {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("CDK_DEFAULT_REGION")),
	})
	if err != nil {
		return err
	}

	svc := dynamodb.New(sess)
	table_name := os.Getenv("VEHICLE_TABLE_ARN")
	// Create the input for the GetItem operation
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(table_name),
		Key: map[string]*dynamodb.AttributeValue{
			"reg_no": {
				S: aws.String(vehicle.RegNo),
			},
		},
	}

	// Get the item from DynamoDB
	result, err := svc.GetItem(getItemInput)
	if err != nil {
		return err
	}

	// Check if the item exists
	if result.Item == nil {
		// Item doesn't exist, insert a new item
		av, err := dynamodbattribute.MarshalMap(vehicle)
		if err != nil {
			return err
		}

		putItemInput := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(table_name),
		}

		_, err = svc.PutItem(putItemInput)
		if err != nil {
			return err
		}

		fmt.Println("New vehicle record inserted successfully")
	} else {
		// Item exists, update the last_check_date
		currentDate := time.Now().Format("02-01-2006")
		updateInput := &dynamodb.UpdateItemInput{
			TableName: aws.String(table_name),
			Key: map[string]*dynamodb.AttributeValue{
				"reg_no": {
					S: aws.String(vehicle.RegNo),
				},
			},
			UpdateExpression: aws.String("SET last_check_date = :cd"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":cd": {
					S: aws.String(currentDate),
				},
			},
		}

		_, err = svc.UpdateItem(updateInput)
		if err != nil {
			return err
		}

		fmt.Println("Vehicle record updated successfully")
	}

	return nil
}
