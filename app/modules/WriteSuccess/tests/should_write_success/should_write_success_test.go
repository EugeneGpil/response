package should_write_success

import (
	"fmt"
	"net/http"
	"testing"

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

	tester.AssertNil(err)

	tester.AssertSame(http.StatusOK, writer.GetStatus())

	messages := writer.GetMessages()

	tester.AssertSame(1, len(messages))

	responseBody := messages[0]

	tester.AssertSame([]byte(expectedResponseBody), responseBody)
}
