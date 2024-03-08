package should_write_message

import (
	"testing"

	"github.com/EugeneGpil/response/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	responsePackage "github.com/EugeneGpil/response"
)

var testMessage = "test message"
var expectedResponseBody = "test message"

func Test_should_write_message(t *testing.T) {
	tester.SetTester(t)

	writer := httpTester.GetTestResponseWriter()

	response := responsePackage.New(writer)

	_, err := response.Write([]byte(testMessage))

	tests.AssertResponse(err, writer, 0, expectedResponseBody)
}
