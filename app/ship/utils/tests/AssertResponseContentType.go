package tests

import (
	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"
)

func AssertResponseContentType(
	writer responseWriter.ResponseWriter,
	expectedResponseContentType string,
) {
	tester.AssertSame(expectedResponseContentType, writer.Header().Get("Content-Type"))
}
