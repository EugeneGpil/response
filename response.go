package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	writer http.ResponseWriter
}

func New(writer http.ResponseWriter) Response {
	return Response{
		writer: writer,
	}
}

func (response Response) WriteValidationErrors(errors map[string]string) error {
	message := struct {
		Message string
		Errors  map[string]string
	}{
		Message: "Your request encountered some errors",
		Errors:  errors,
	}

	responseBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	response.writer.Write(responseBody)
	response.writer.WriteHeader(http.StatusUnprocessableEntity)

	return nil
}
