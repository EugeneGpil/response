package WriteValidationErrors

import (
	"net/http"

	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors/interfaces"
	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors/vars"

	shipWriter "github.com/EugeneGpil/response/app/ship/utils/writer"
)

func WriteValidationErrors(writer http.ResponseWriter, dto interfaces.ValidationErrorsDto) error {
	message := getMessage(dto.GetMessage())

	body := struct {
		Message string
		Errors  map[string]string
		Code    int
	}{
		Message: message,
		Errors:  dto.GetErrors(),
		Code:    dto.GetCode(),
	}

	return shipWriter.
		New().
		SetWriter(writer).
		Write(http.StatusUnprocessableEntity, body)
}

func getMessage(message string) string {
	if message == "" {
		return vars.DefaultMessage
	}

	return message
}
