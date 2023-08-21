package main

import (
	"os"

	"C:\Users\andrecristhian.barre\OneDrive - S3K S.p.A\Documents\gitprojects\CRUD-API-GO\pkg\handlers"
	"pkg\handlers"
	"github.com/aws/aws-sdk-go/lambda/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbiface"
	"github.com/aws/aws-sdk-go/lambda"
)

var( dynaClient dynamodbiface.DynamoDBAPI )

func main(){
	region:= os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{region:aws.String(region)},) 

	if err !=nil{
		return
	}

	dynaClient := dynamodb.New(awsSession)
	lambda.Start(handler)

	}

	const tableName = "LambdaInGoUser"

	func handler(req events.APIGatewayProxyRequest)(*events.APIGatewayProxyResponse, error){
		switch req.HTTPMethod{
		case "GET":
			return handlers.GetUser(req, tableName, dynaClient)
		case "POST":
			return handlers.CreateUser(req, tableName, dynaClient)
		case "PUT":
			return handlers.UpdateUser(req, tableName, dynaClient)
		case "DELETE":
			return handlers.DeleteUser(req, tableName, dynaClient)
		default:
			return handlers.UnhandleMethod()
	}