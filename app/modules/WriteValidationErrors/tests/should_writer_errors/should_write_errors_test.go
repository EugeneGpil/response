package should_write_errors

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/EugeneGpil/response/app/modules/WriteValidationErrors/vars"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	responsePackage "github.com/EugeneGpil/response"
)

var property = "property"
var responseError = "some error"
var expectedMessage = fmt.Sprintf(
	"{\"Message\":\"%s\",\"Errors\":{\"%s\":\"%s\"},\"Code\":0}",
	vars.DefaultMessage,
	property,
	responseError,
)

func Test_should_write_errors(t *testing.T) {
	tester.SetTester(t)

	errors := map[string]string{
		property: responseError,
	}

	writer := httpTester.GetTestResponseWriter()

	response := responsePackage.New(writer)

	err := response.WriteValidationErrors(responsePackage.ValidationErrorsDto{
		Errors: errors,
	})

	tester.AssertNil(err)

	tester.AssertSame(http.StatusUnprocessableEntity, writer.GetStatus())

	messages := writer.GetMessages()

	tester.AssertSame(1, len(messages))

	message := messages[0]

	tester.AssertSame([]byte(expectedMessage), message)
}
