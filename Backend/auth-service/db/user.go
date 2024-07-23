package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type User struct {
	OfficeName string `dynamodbav:"office_name"`
	Email      string `dynamodbav:"email"`
	Password   string `dynamodbav:"password"`
}

func CreateUserByEmail(u *User) (*User, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	svc := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("USER_TABLE_ARN")
	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		log.Println("Error marshaling user struct to AttributeValue map")
		return nil, err
	}
	condition := expression.AttributeNotExists(expression.Name("email"))
	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		log.Println("Error building expression")
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:                      av,
		TableName:                 aws.String(tableName),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}
	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByEmail(Email string) (*User, error) {
	var users []*User
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("USER_TABLE_ARN")

	keyCond := expression.Key("email").Equal(expression.Value(Email))
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
		log.Println("error querying items : ", err)
		return nil, err
	}
	err = attributevalue.UnmarshalListOfMaps(output.Items, &users)
	if err != nil {
		log.Println("error unmarshalling attributes : ", err)
		return nil, err
	}

	//* Check if there are no offices registered with the given email address
	if len(users) == 0 {
		log.Println("NO Authorities")
		return &User{}, fmt.Errorf("no authorities registered with the given email address")
	}

	return users[0], nil
}
