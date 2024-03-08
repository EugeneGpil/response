package tests

import "github.com/EugeneGpil/responseWriter"

func AssertResponseContentTypeApplicationJson(writer responseWriter.ResponseWriter) {
	AssertResponseContentType(writer, "application/json")
}
