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
