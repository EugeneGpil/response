package WriteValidationErrors

import (
	"encoding/json"
	"net/http"

	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors/interfaces"
	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors/vars"
)

func WriteValidationErrors(writer http.ResponseWriter, dto interfaces.ValidationErrorsDto) error {
	message := getMessage(dto.GetMessage())

	responseBodyStruct := struct {
		Message string
		Errors  map[string]string
		Code    int
	}{
		Message: message,
		Errors:  dto.GetErrors(),
		Code:    dto.GetCode(),
	}

	responseBody, err := json.Marshal(responseBodyStruct)
	if err != nil {
		return err
	}

	writer.WriteHeader(http.StatusUnprocessableEntity)
	writer.Write(responseBody)

	return nil
}

func getMessage(message string) string {
	if message == "" {
		return vars.DefaultMessage
	}

	return message
}
