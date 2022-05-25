package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var buf []byte
	var (
		result string
		err    error
	)
	ch.Input.Read(buf)
	result, err = CalculatePostfix(string(buf))
	if result == "Nil" && err != nil {
		ch.Output.Write([]byte(err.Error()))
	} else if result != "Nil" {
		ch.Output.Write([]byte(result))
	} else {
		return err
	}
	return nil
}
