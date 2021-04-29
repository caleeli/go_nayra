package aws_common

import (
	"encoding/json"
	"fmt"
)

type ResponseHttp struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

func SuccessfulResponse(response interface{}) (ResponseHttp, error) {
	dataJson, _ := json.Marshal(response)
	return ResponseHttp{
		StatusCode: 200,
		Headers: map[string]string{
			"access-control-allow-origin":  "*",
			"access-control-allow-headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"access-control-allow-methods": "OPTIONS,GET,POST",
		},
		Body:            string(dataJson),
		IsBase64Encoded: false,
	}, nil
}

func ErrorResponse(err error, message string, args ...interface{}) (ResponseHttp, error) {
	dataJson, _ := json.Marshal(struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: false,
		Message: fmt.Sprintf(message, args...),
	})
	return ResponseHttp{
		StatusCode: 500,
		Headers: map[string]string{
			"access-control-allow-origin":  "*",
			"access-control-allow-headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"access-control-allow-methods": "OPTIONS,GET,POST",
		},
		Body:            string(dataJson),
		IsBase64Encoded: false,
	}, err
}
