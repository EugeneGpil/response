package should_write_success

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/EugeneGpil/response/app/ship/utils/tests"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	responsePackage "github.com/EugeneGpil/response"
)

var testMessage = "test message"
var expectedResponseBody = fmt.Sprintf(`{"Message":"%v"}`, testMessage)

func Test_should_write_success(t *testing.T) {
	tester.SetTester(t)

	writer := httpTester.GetTestResponseWriter()

	response := responsePackage.New(writer)

	err := response.WriteSuccess(testMessage)

	tests.AssertResponse(err, writer, http.StatusOK, expectedResponseBody)
}
