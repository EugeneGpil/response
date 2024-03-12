package response

import (
	"net/http"

	"github.com/EugeneGpil/response/app/modules/WriteSuccess"
	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors"
)

type Response struct {
	writer http.ResponseWriter
}

func New(writer http.ResponseWriter) Response {
	return Response{
		writer: writer,
	}
}

func (response Response) WriteValidationErrors(dto ValidationErrorsDto) error {
	return WriteValidationErrors.WriteValidationErrors(response.writer, dto)
}

func (response Response) WriteSuccess(message string) error {
	return WriteSuccess.WriteSuccess(response.writer, message)
}

func (response Response) Write(message []byte) (int, error) {
	return response.writer.Write(message)
}

func (response Response) SetStatusCode(statusCode int) {
	response.writer.WriteHeader(statusCode)
}

func (response Response) WriteBadRequest(error string) {
	http.Error(response.writer, error, http.StatusBadRequest)
}
