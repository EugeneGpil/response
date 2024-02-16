package writer

import (
	"encoding/json"
	"net/http"
)

type writer struct {
	writer http.ResponseWriter;
	status int;
}

func New() *writer {
	return &writer{
		status: http.StatusOK,
	}
}

func (shipWriter *writer) SetWriter(writer http.ResponseWriter) *writer {
	shipWriter.writer = writer

	return shipWriter
}

func (shipWriter *writer) Write(status int, body interface{}) error {
	responseBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	shipWriter.writer.Header().Add("Content-Type", "application/json")
	shipWriter.writer.WriteHeader(status)
	shipWriter.writer.Write(responseBody)

	return nil
}
