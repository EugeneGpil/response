package should_write_errors_with_message_and_code

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	responsePackage "github.com/EugeneGpil/response"
)

var mainErrorMessage = "Main error message"
var property = "property"
var responseError = "some error"
var errorCode = 15
var expectedResponseBody = fmt.Sprintf(
	`{"Message":"%v","Errors":{"%v":"%v"},"Code":%v}`,
	mainErrorMessage,
	property,
	responseError,
	errorCode,
)

func Test_should_write_errors_with_message_and_code(t *testing.T) {
	tester.SetTester(t)

	errors := map[string]string{
		property: responseError,
	}

	writer := httpTester.GetTestResponseWriter()

	response := responsePackage.New(writer)

	err := response.WriteValidationErrors(responsePackage.ValidationErrorsDto{
		Message: mainErrorMessage,
		Errors:  errors,
		Code:    errorCode,
	})

	tester.AssertNil(err)

	tester.AssertSame(http.StatusUnprocessableEntity, writer.GetStatus())

	messages := writer.GetMessages()

	tester.AssertSame(1, len(messages))

	responseBody := messages[0]

	tester.AssertSame([]byte(expectedResponseBody), responseBody)
}
