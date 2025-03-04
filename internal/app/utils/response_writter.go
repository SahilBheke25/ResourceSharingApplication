package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleResponse(w http.ResponseWriter, message any, r *http.Request) {

	res, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type response struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

func SuccessResponse(ctx context.Context, w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	payload := response{
		Data: data,
	}

	out, err := json.Marshal(payload)
	if err != nil {
		log.Println(ctx, "cannot marshal success response payload, err : ", err)
		writeServerErrorResponse(ctx, w)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Println(ctx, "cannot write json success response, err : ", err)
		writeServerErrorResponse(ctx, w)
		return
	}
}

func ErrorResponse(ctx context.Context, w http.ResponseWriter, httpStatus int, err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	payload := response{
		ErrorCode:    httpStatus,
		ErrorMessage: err.Error(),
	}
	out, err := json.Marshal(payload)
	if err != nil {
		log.Println(ctx, "error occured while marshaling response payload, err : ", err)
		writeServerErrorResponse(ctx, w)
		return
	}

	_, err = w.Write(out)

	if err != nil {
		log.Println(ctx, "error occured while writing response, err : ", err)
		writeServerErrorResponse(ctx, w)
		return
	}
}

func writeServerErrorResponse(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "internal server error")))
	if err != nil {
		log.Println(ctx, "error occured while writing response, err : ", err)
	}
}
