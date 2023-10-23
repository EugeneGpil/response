package WriteSuccess

import (
	"encoding/json"
	"net/http"
)

func WriteSuccess(writer http.ResponseWriter, message string) error {
	responseBodyStruct := struct{
		Message string
	} {
		Message: message,
	}

	responseBody, err := json.Marshal(responseBodyStruct)
	if err != nil {
		return err
	}

	writer.Write(responseBody)
	writer.WriteHeader(http.StatusOK)

	return nil
}
