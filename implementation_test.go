package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCalculatePostfix(c *C) {
	var res string
	var err error
	res, err = CalculatePostfix("1 2 + 6 - 2 *") // (1 + 2 - 6) * 2
	c.Assert(res, Equals, "-6.000000")
	c.Assert(err, IsNil)
	res, err = CalculatePostfix("1 4 + -8.5 -") // 1 + 4 - (-8.5)
	c.Assert(res, Equals, "13.500000")
	c.Assert(err, IsNil)
	res, err = CalculatePostfix("1 10 + -2 / 3 2 ^ +") // (1 + 10)/(-2) + 3^2
	c.Assert(res, Equals, "3.500000")
	c.Assert(err, IsNil)
	res, err = CalculatePostfix("35 13 * 2 / 310 + 2 * 33 + 22.4 - 2 ^") // ((35 * 13 / 2 + 310) * 2 + 33 - 22.4)^2
	c.Assert(res, Equals, "1178527.360000")
	c.Assert(err, IsNil)
	res, err = CalculatePostfix("9 1 / 3 + 2 + 6 - 3 * 63 - 22 - -3.4 / 5.3 ^") // (((9 / 1 + 3 + 2 - 6) * 3 - 63 - 22) / -3.4) ^ 5.3
	c.Assert(res, Equals, "4419873.759829")
	c.Assert(err, IsNil)
	res, err = CalculatePostfix("5 -5 * 3 * 4 - 3 + -12.4 - 55 + 3 ^ 3 / 15335 - 33 /") // ((5 * -5 * 3 - 4 + 3 - -12.4 + 55) ^ 3 / 3 - 15335) / 33
	c.Assert(res, Equals, "-471.121778")
	c.Assert(err, IsNil)
	res, err = CalculatePostfix("")
	c.Assert(res, Equals, "Nil")
	c.Assert(err, ErrorMatches, "Wrong expression")
	res, err = CalculatePostfix("qwerty")
	c.Assert(res, Equals, "Nil")
	c.Assert(err, ErrorMatches, "Wrong value.+")
	res, err = CalculatePostfix("1 2 3")
	c.Assert(res, Equals, "Nil")
	c.Assert(err, ErrorMatches, "Wrong expression")
	res, err = CalculatePostfix("+ -")
	c.Assert(res, Equals, "Nil")
	c.Assert(err, ErrorMatches, "There are no 2 values, can't calculate expression.+")
	res, err = CalculatePostfix("2 3 + -")
	c.Assert(res, Equals, "Nil")
	c.Assert(err, ErrorMatches, "There are no 2 values, can't calculate expression.+")
}

func ExampleCalculatePostfix() {
	res, err := CalculatePostfix("3 5 + 2 /")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// 4.000000
}
