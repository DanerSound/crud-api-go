package handlers

import(
	"encoding/json"
	"github.com/aws/aws-sdk-go/events"
)

func apiResponse(status int, body interface{})(*events.APIGatewayProxyResponse, error){
	resp:= events.APIGatewayProxyResponse{Headers: map[string]string["Content-Type":"application/json"]}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body) // because GOLANG dont undertand json, need this step to managed json
	resp.Body = string(stringBody)

}