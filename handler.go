package lab2

import (
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var buf = make([]byte, 20)
	var (
		result string
		err    error
		exp    string
	)
	ch.Input.Read(buf)
	exp = string(buf)
	result, err = CalculatePostfix(exp)
	if result == "Nil" && err != nil {
		ch.Output.Write([]byte(err.Error()))
	} else if result != "Nil" {
		fmt.Println(result)
		ch.Output.Write([]byte(result))
	} else {
		return err
	}
	return nil
}
