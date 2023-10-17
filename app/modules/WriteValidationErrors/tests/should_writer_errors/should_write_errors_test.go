package should_write_errors

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

var property = "property"
var responseError = "some error"
var expectedMessage = "{\"Message\":\"Your request encountered some errors\",\"Errors\":{\"" + property + "\":\"" + responseError + "\"}}"

func Test_should_write_errors(t *testing.T) {
	tester.SetTester(t)

	errors := map[string]string{
		property: responseError,
	}

	writer := httpTester.GetTestResponseWriter()

	response := response.New(writer)

	err := response.WriteValidationErrors(errors)

	tester.AssertNil(err)

	tester.AssertSame(http.StatusUnprocessableEntity, writer.GetStatus())

	messages := writer.GetMessages()

	tester.AssertSame(1, len(messages))

	message := messages[0]

	tester.AssertSame([]byte(expectedMessage), message)
}
