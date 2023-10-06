package dynamodbgo

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/marcosmourabro/dynamo-data-dispatcher/model"
)

func CreateDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)
	return svc
}

func InsertItem(connection *dynamodb.DynamoDB, tablename string) (events.APIGatewayProxyResponse, error) {

	var product model.Product

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tablename),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(product.ID),
			},
			"name": {
				S: aws.String(product.NAME),
			},
			"price": {
				N: aws.String(strconv.Itoa(int(product.PRICE))),
			},
		},
	}

	_, err := connection.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	body, err := json.Marshal(product)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
