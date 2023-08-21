package handlers

import (
	"crud-api-go/pkg/user"
	"net/http"
	"pkg/user"

	"github.com/aws/aws-lambda-go/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

var ErrorMethodNotAllows = "Method not allowed"

type ErrorBody struct{
	ErrorMsg *string `json: "error, omitempty"`
}


func GetUser(req events.APIGatewayProxyRequest, tableName string dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error) {

	email := req.QueryStringParamters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err!= nil{
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(email, tableName, dynaClient)
	if err!= nil{
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error()),})
		}
		return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error) {

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error) {

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatewayProxyResponse, error) {}

func UnhandleMethod()(*events.APIGatewayProxyResponse, error) {

	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllows)
}
