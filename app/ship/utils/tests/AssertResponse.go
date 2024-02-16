package tests

import (
	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"
)

func AssertResponse(
	err error,
	writer responseWriter.ResponseWriter,
	expectedStatusCode int,
	expectedResponseBody string,
) {
	tester.AssertNil(err)

	tester.AssertSame(expectedStatusCode, writer.GetStatus())

	tester.AssertSame("application/json", writer.Header().Get("Content/type"))

	messages := writer.GetMessages()

	tester.AssertSame(1, len(messages))

	responseBody := messages[0]

	tester.AssertSame([]byte(expectedResponseBody), responseBody)
}
