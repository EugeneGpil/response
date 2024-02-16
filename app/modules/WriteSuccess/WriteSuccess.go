package WriteSuccess

import (
	"net/http"

	shipWriter "github.com/EugeneGpil/response/app/ship/utils/writer"
)

func WriteSuccess(writer http.ResponseWriter, message string) error {
	body := struct{
		Message string
	} {
		Message: message,
	}

	return shipWriter.WriteResponse(writer, http.StatusOK, body)
}
