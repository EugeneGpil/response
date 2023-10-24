package should_write_errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors/vars"
	"github.com/EugeneGpil/responseWriter"
	"github.com/EugeneGpil/tester"

	
	responsePackage "github.com/EugeneGpil/response"
)

var property = "property"
var responseError = "some error"
var expectedResponseBody = fmt.Sprintf(
	`{"Message":"%v","Errors":{"%v":"%v"},"Code":0}`,
	vars.DefaultMessage,
	property,
	responseError,
)

func Test_should_write_errors(t *testing.T) {
	tester.SetTester(t)

	errors := map[string]string{
		property: responseError,
	}

	writer := responseWriter.New()

	response := responsePackage.New(writer)

	err := response.WriteValidationErrors(responsePackage.ValidationErrorsDto{
		Errors: errors,
	})

	tester.AssertNil(err)

	tester.AssertSame(http.StatusUnprocessableEntity, writer.GetStatus())

	messages := writer.GetMessages()

	tester.AssertSame(1, len(messages))

	responseBody := messages[0]

	tester.AssertSame([]byte(expectedResponseBody), responseBody)
}
