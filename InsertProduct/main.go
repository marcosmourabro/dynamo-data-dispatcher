package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	dynamodbgo "github.com/marcosmourabro/dynamo-data-dispatcher/dynamodb.go"
)

func main() {

	connection := dynamodbgo.CreateDynamoSession()
	InsertItem, _ := dynamodbgo.InsertItem(connection, "itens")
	lambda.Start(InsertItem)
}
