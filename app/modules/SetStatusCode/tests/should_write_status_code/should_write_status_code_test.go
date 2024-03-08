package should_write_status_code

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
	responsePackage "github.com/EugeneGpil/response"
)

func Test_should_write_status_code(t *testing.T) {
	tester.SetTester(t)

	writer := httpTester.GetTestResponseWriter()

	tester.AssertSame(0, writer.GetStatus())

	response := responsePackage.New(writer)

	response.SetStatusCode(http.StatusUnprocessableEntity)

	tester.AssertSame(http.StatusUnprocessableEntity, writer.GetStatus())
}
