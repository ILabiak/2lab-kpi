package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/ILabiak/2lab-kpi"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File to take expression from")
	outputFile      = flag.String("o", "", "Output file")
	handler         lab2.ComputeHandler
	reader          io.Reader = nil
	writer          io.Writer = nil
)

func main() {
	flag.Parse()
	if *inputExpression == "" && *inputFile == "" {
		panic("No expression specified")
	} else if *inputExpression != "" {
		reader = strings.NewReader(strings.Trim(*inputExpression, " "))
	} else if *inputFile != "" {
		var (
			exp []byte
			err error
		)

		exp, err = os.ReadFile(*inputFile)
		if err == nil {
			reader = strings.NewReader(string(exp))
		} else {
			panic(err)
		}
	}

	if *outputFile != "" {
		var err error
		writer, err = os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
	} else {
		writer = os.Stdout
	}
	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	var err error = handler.Compute()
	if err != nil {
		panic(err)
	}
}
