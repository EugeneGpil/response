package writer

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(writer http.ResponseWriter, status int, body interface{}) error {
	responseBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(responseBody)

	return nil
}
