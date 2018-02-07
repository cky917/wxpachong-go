package spec

import (
	"encoding/json"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func MarshalResponse(data interface{}, err error) (result []byte, statusCode int) {
	resp := Response{0, "ok", data}
	statusCode = 200
	if err != nil {
		resp.Code, resp.Message = HandleError(err)
		resp.Data = nil
		statusCode = resp.Code
	}
	result, err = json.Marshal(resp)
	if err != nil {
		statusCode = 599
		result = []byte(err.Error())
	}
	return
}
