package lab2

import (
	"os"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func HandlerTest(t *testing.T) { TestingT(t) }

type MyHandlerTestSuite struct{}

var _ = Suite(&MyHandlerTestSuite{})

func (s *MyHandlerTestSuite) TestHandlerCompute(c *C) {
	var (
		res        string
		err        error
		inputError error
	)

	handler := &ComputeHandler{
		Input:  strings.NewReader("2 5 ^"),
		Output: os.Stdout,
	}
	err = handler.Compute()
	res, inputError = CalculatePostfix("2 5 ^")
	c.Assert(inputError, IsNil)
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "32.000000")
	handler = &ComputeHandler{
		Input:  strings.NewReader("wrong input"),
		Output: os.Stdout,
	}
	err = handler.Compute()
	c.Assert(err, ErrorMatches, "Wrong value - \"wrong\"")
}
