package should_write_errors_with_message_and_code

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/EugeneGpil/response/app/ship/utils/tests"
	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"

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

	writer := responseWriter.New()

	response := responsePackage.New(writer)

	err := response.WriteValidationErrors(responsePackage.ValidationErrorsDto{
		Message: mainErrorMessage,
		Errors:  errors,
		Code:    errorCode,
	})

	tests.AssertResponse(err, writer, http.StatusUnprocessableEntity, expectedResponseBody)
}
