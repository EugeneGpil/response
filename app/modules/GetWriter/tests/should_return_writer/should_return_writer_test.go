package should_return_writer

import (
	"testing"

	"github.com/EugeneGpil/response"
	"github.com/EugeneGpil/tester"

	httpTester "github.com/EugeneGpil/httpTester"
)

func Test_should_return_writer(t *testing.T) {
	tester.SetTester(t)

	writer := httpTester.GetTestResponseWriter()

	response := response.New(writer)

	actualWriter := response.GetWriter()

	tester.AssertSame(writer, actualWriter)
}
