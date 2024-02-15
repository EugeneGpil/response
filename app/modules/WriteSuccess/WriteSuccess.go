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

	writer.WriteHeader(http.StatusOK)
	writer.Write(responseBody)

	return nil
}
