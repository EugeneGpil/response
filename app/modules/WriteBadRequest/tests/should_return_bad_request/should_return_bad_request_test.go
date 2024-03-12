package should_return_bad_request

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"
)

func Test_should_return_bad_request(t *testing.T) {
	tester.SetTester(t)

	writer := responseWriter.New()

	response := response.New(writer)

	error := "error"

	response.WriteBadRequest(error)

	tester.AssertSame(http.StatusBadRequest, writer.GetStatus())

	contentType := writer.Header().Get("Content-Type")

	tester.AssertSame("text/plain; charset=utf-8", contentType)

	XContentTypeOptions := writer.Header().Get("X-Content-Type-Options")

	tester.AssertSame("nosniff", XContentTypeOptions)
}
