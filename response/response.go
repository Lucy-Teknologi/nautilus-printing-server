package response

import (
	log "nautilus-print-server/log"

	jsoniter "github.com/json-iterator/go"
)

type Status string

type Response struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Response) ToByte() []byte {
	bytes, err := jsoniter.Marshal(r)
	if err != nil {
		log.Default().Printf("Error marshalling response: %s", err)
	}
	return bytes
}

func Success(data interface{}) Response {
	return Response{
		Status:  OK,
		Message: "Successfully ran operation",
		Data:    data,
	}
}

func Error(err error) Response {
	return Response{
		Status:  ERROR,
		Message: err.Error(),
	}
}

func ErrorWithMessage(message string) Response {
	return Response{
		Status:  ERROR,
		Message: message,
	}
}
